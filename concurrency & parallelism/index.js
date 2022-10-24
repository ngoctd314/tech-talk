let displaySection = 1;

function decrement() {
    let newSection = displaySection
    if (displaySection >1 ) {
        newSection = displaySection-1;
    }
    if (newSection != displaySection) {
        document.getElementById(displaySection).setAttribute("style", "display:none")
        document.getElementById(newSection).setAttribute("style", "")
        displaySection = newSection;
    }
}

function increment() {
    let newSection = displaySection
    if (displaySection < 2 ) {
        newSection = displaySection+1;
    }
    if (newSection != displaySection) {
        document.getElementById(displaySection).setAttribute("style", "display:none")
        document.getElementById(newSection).setAttribute("style", "")
        displaySection = newSection;
    }
}

let prevBtn = document.getElementById("prev_btn");
prevBtn.addEventListener("click", decrement, false);

let nextBtn = document.getElementById("next_btn");
nextBtn.addEventListener("click", increment, false);