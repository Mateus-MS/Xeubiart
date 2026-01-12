// TODO: make it re-usable when necessary
class CustomSelect{
    constructor(containerID){
        this.container = document.getElementById(containerID)
        this.optsHolder = this.container.querySelector("ul")
        this.opts = this.optsHolder.children
        this.selected = this.container.querySelector(".selected")
        this.hightlitedIndex = -1
        
        this.registerEventListeners()
    }

    ToggleOpenState(){
        this.optsHolder.classList.toggle("hidden")
    }

    SelectOption(index){
        this.selected.innerText = this.opts[index].innerText
    }

    HighlightOption(index){
        if(this.hightlitedIndex >= 0){
            // un-highlight the last one
            this.opts[this.hightlitedIndex].classList.remove("highlight")
        }

        // Now uses the new value
        this.hightlitedIndex = index

        // And highlight the new one
        this.opts[this.hightlitedIndex].classList.add("highlight")
    }

    registerEventListeners(){
        this.selected.addEventListener("click", (e)=>{
            this.ToggleOpenState()
        })

        // Add a listener to each opt
        for(let i = 0; i < this.opts.length; i++){
            this.opts[i].addEventListener("click", (e)=>{
                this.SelectOption(i)
                this.ToggleOpenState()
            })

            this.opts[i].addEventListener("mouseenter", (e)=>{
                if(this.optsHolder.classList.contains("hidden")) return

                this.HighlightOption(i)
            })
        }

        // Accessibility for desktop
        window.addEventListener("keydown", (e)=>{
            if(this.optsHolder.classList.contains("hidden")){
                return
            }

            if (e.key === "ArrowDown") {
                if(this.hightlitedIndex + 1 < this.opts.length){
                    this.HighlightOption(this.hightlitedIndex+1)
                }
            }
            
            if (e.key === "ArrowUp") {
                if(this.hightlitedIndex - 1 >= 0){
                    this.HighlightOption(this.hightlitedIndex-1)
                }
            }
        })
    }

}

new CustomSelect("custom-select")