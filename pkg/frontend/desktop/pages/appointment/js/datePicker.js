var monthDaysContainer = document.getElementById("date_picker__montDays__container")

async function updateSchedule(){
    let data = await requestScheduleData(2026, 1)

    let firstWeekDay = parseInt(data.date.firstWeekday)
    let daysInMonth = parseInt(data.date.daysInMonth)

    let appointments = data.schedule.appointments

    await loadDays(firstWeekDay, daysInMonth)
    await loadAppointments(appointments, firstWeekDay)
}

async function loadDays(firstWeekDay, daysInMonth){
    let children = monthDaysContainer.children
    let dayCounter = 1

    for(let i = 0; i < children.length; i++){
        if(i < firstWeekDay){
            continue
        }

        if(dayCounter > daysInMonth){
            continue
        }

        children[i].innerText = dayCounter
        children[i].dataset.day = dayCounter
        dayCounter += 1
    }
}

async function loadAppointments(appointments, firstWeekDay){
    for(let i = 0; i < appointments.length; i++){
        let date = appointments[i].Date.split("T")[0]
        let [year, month, day] = date.split("-").map(Number)

        monthDaysContainer.querySelector(`[data-day="${day}"]`).classList.add("date_picker__monthDay-appointed")
    }
}

async function requestScheduleData(year, month) {
    try {
        let response = await fetch(`http://localhost:8080/api/schedule/${year}/${month}`, {method: "GET"})
        if (!response.ok) {
            throw new Error(`HTTP error ${response.status}`);
        }

        let data = await response.json();
        return data
    } catch (err) {
        console.error(err);
    }
}