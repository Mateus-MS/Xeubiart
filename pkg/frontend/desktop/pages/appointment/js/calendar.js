class HourOption{
    constructor(hourOptionElement){
        this.element = hourOptionElement

        this.initiateEventListeners()
    }

    highlight(){
        this.element.classList.add("highlight")
    }
    
    unHighlight(){
        this.element.classList.remove("highlight")
    }

    show(){
        this.element.classList.remove("hidden")
    }

    hide(){
        this.element.classList.add("hidden")
    }

    getText(){
        return this.element.innerText
    }

    onClick(){
        console.log("clicked")
    }

    initiateEventListeners(){
        this.element.addEventListener("mouseenter", ()=>{
            // If the parent is on "closing" state, don't highlight
            if(this.element.parentNode.classList.contains("hidden")) return
            this.highlight()
        })

        this.element.addEventListener("mouseout", ()=>{
            this.unHighlight()
        })

        this.element.addEventListener("click", ()=>{
            this.onClick()
        })
    }
}

class HourDisplay{
    constructor(hourDisplayElement){
        this.element = hourDisplayElement
        this.optsHolder = this.element.parentNode.querySelector("ul")

        this.initiateEventListeners()
    }

    toggleOpenState(){
        this.optsHolder.classList.toggle("hidden")
    }

    initiateEventListeners(){
        this.element.addEventListener("click", ()=>{
            this.toggleOpenState()
        })
    }

    updateText(newText){
        this.element.innerText = newText
    }
}

class HourSelector{
    constructor(){
        this.hoursOptions = [];
        for(let element of document.getElementById("custom-hours-holder").children){
            this.hoursOptions.push(new HourOption(element))
        }

        this.selectedDisplay = new HourDisplay(document.getElementById("custom-selected-display"))
        this.hightlitedIndex = null

        this.initiateEventListeners()
    }

    getSelectedHour(){
        return this.selectedDisplay.element.innerText
    }

    reset(){
        for(let opt of this.hoursOptions){
            opt.show()
        }
    }

    // startHour should be in the request format '2026-01-15T15:00:00Z'
    hideHours(startHour, hoursToHide){
        startHour = startHour.split("T")[1] //'15:00:00Z'
        startHour = startHour.slice(0, -4)  //'15:00'

        for(let i = 0; i < this.hoursOptions.length; i++){
            if(startHour === this.hoursOptions[i].getText()){
                for(let j = 0; j < hoursToHide * 2; j++){
                    if(this.hoursOptions[i + j]){
                        this.hoursOptions[i + j].hide()
                    }
                }
                return
            }
        }
    }

    initiateEventListeners(){
        for(let i = 0; i < this.hoursOptions.length; i++){
            let child = this.hoursOptions[i]
            child.onClick = ()=>{
                child.highlight()
                this.selectedDisplay.toggleOpenState()
                this.selectedDisplay.updateText(child.getText())
            }
        }
    }
}

class Calendar {
    constructor(){
        this.hourSelector = new HourSelector()

        // Elements
        this.daysContainer = document.getElementById("monthDays")
        this.prevButtonElement = document.getElementById("calendar-prev")
        this.nextButtonElement = document.getElementById("calendar-next")
        this.makeAppointmentButton = document.getElementById("make-appointment")
        this.monthYearDisplay = document.getElementById("date-display")

        this.monthOffset = 0
        this.selectedDay = null

        // Fetched data
        this.year         = null
        this.firstWeekDay = null
        this.daysInMonth  = null
        this.today        = null
        this.bookedHours  = null
        this.actualMonth  = null

        this.init()
    }

    async update(){
        this.reset()

        await this.syncData()

        this.createDays()

        // Update the month/year display
        this.monthYearDisplay.innerText = `${convertIntToMonthName(this.actualMonth)} ${this.year}`
    }

    async syncData(){
        try {
            let response = await fetch(`http://localhost:8080/api/schedule/${this.monthOffset}`, {method: "GET"})
            if (!response.ok) {
                throw new Error(`HTTP error ${response.status}`);
            }

            let data = await response.json();

            this.firstWeekDay = parseInt(data.date.firstWeekday)
            this.daysInMonth  = parseInt(data.date.daysInMonth)
            this.today        = parseInt(data.date.today)
            this.bookedHours  = new Map(Object.entries(data.schedule.days))
            this.selectedDay  = this.today
            this.year         = parseInt(data.date.year)
            this.actualMonth  = parseInt(data.date.month)
        } catch (err) {
            console.error(err);
        }
    }

    getSelectedDay(){
        return this.daysContainer.querySelector(`[data-day="${this.selectedDay}"]`).innerText
    }

    // Uses the fetched data to initiate the calendar days
    createDays(){
        if(this.firstWeekDay == null) return
        if(this.daysInMonth == null) return
        if(this.today == null) return
        if(this.bookedHours == null) return

        let children = this.daysContainer.children
        let dayCounter = 1

        for(let i = 0; i < children.length; i++){
            if(i < this.firstWeekDay || dayCounter > this.daysInMonth){
                continue
            }
            
            if(dayCounter <= this.today){
                children[i].classList.add("pastDay")
            }

            if(dayCounter === this.today){
                children[i].classList.add("selected")
            }
            
            if(dayCounter >= this.today && dayCounter <= this.daysInMonth){
                children[i].classList.add("available")
            }

            children[i].innerText = dayCounter
            children[i].dataset.day = dayCounter
            dayCounter += 1
        }
    }

    highlightDay(index){
        this.daysContainer.querySelector(`[data-day="${this.selectedDay}"]`).classList.remove("selected")
        this.selectedDay = index
        this.daysContainer.querySelector(`[data-day="${this.selectedDay}"]`).classList.add("selected")
    }

    reset(){
        for(let child of this.daysContainer.children){
            child.classList.remove("pastDay")
            child.classList.remove("selected")
            child.classList.remove("available")
            child.innerText = ""
            child.dataset.day = ""
        }

        this.hourSelector.reset()
    }

    init(){
        window.addEventListener("load", (e)=>{
            this.update()
        })

        let children = this.daysContainer.children
        for(let i = 0; i < children.length; i++){
            children[i].addEventListener("click", (e)=>{
                if (children[i].dataset.day === "") {
                    return
                }

                console.log(children[i].dataset.day)

                this.highlightDay(children[i].dataset.day)
                this.hourSelector.reset()

                let appointments = this.bookedHours.get(children[i].dataset.day)
                if(appointments === undefined){
                    return
                }

                for(let appointment of appointments){
                    this.hourSelector.hideHours(appointment, 2)
                }
            })
        }

        this.nextButtonElement.addEventListener("click", (e)=>{
            if(this.monthOffset + 1 > 12){
                return
            }

            this.monthOffset += 1
            this.update()

            if(this.monthOffset == 12){
                this.nextButtonElement.classList.add("hidden")
            }
            
            if(this.monthOffset > 0){
                this.prevButtonElement.classList.remove("hidden")
            }
        })

        this.prevButtonElement.addEventListener("click", (e)=>{
            if(this.monthOffset - 1 < 0){
                return
            }

            this.monthOffset -= 1
            this.update()

            if(this.monthOffset == 0){
                this.prevButtonElement.classList.add("hidden")
            }

            if(this.monthOffset < 12){
                this.nextButtonElement.classList.remove("hidden")
            }
        })

        this.makeAppointmentButton.addEventListener("click", async ()=>{
            try {
                let response = await fetch(`http://localhost:8080/api/appointment`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        userID: "65a1234e9c0f1a2b3c4d5e6f", // TODO: change to a session token, and send it from a cookie
                        date: `${this.year}-${this.actualMonth.length == 2 ? this.actualMonth : "0" + this.actualMonth}-${this.getSelectedDay()} ${this.hourSelector.getSelectedHour()}:00`,
                    })
                })
                
                if (!response.ok) {
                    throw new Error(`HTTP error ${response.status}`);
                }

                if (response.ok) {
                    // TODO: reload the page. (when the page realods, since now the user has a appointment registered, he gonna se a different page)
                }
            } catch (err) {
                console.error(err);
            }
        })
    }
}

function convertIntToMonthName(monthInt){
    switch(monthInt){
        case 1:
            return "January"
        case 2:
            return "February"
        case 3:
            return "March"
        case 4:
            return "April"
        case 5:
            return "May"
        case 6:
            return "June"
        case 7:
            return "July"
        case 8:
            return "August"
        case 9:
            return "September"
        case 10:
            return "October"
        case 11:
            return "November"
        case 12:
            return "December"
    }
}

new Calendar()