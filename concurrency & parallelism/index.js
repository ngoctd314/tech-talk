let displaySection = 1;
let maxSection = 5

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
    if (displaySection < maxSection ) {
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

function swapImg1_2() {
    id = document.getElementById("img-1.2")
    src = id.getAttribute("src")
    if (src == "assets/1.2.svg") {
        id.setAttribute("src", "assets/1.2.1.svg")
    }else {
        id.setAttribute("src", "assets/1.2.svg")
    }
}
document.getElementById("img-1.2").addEventListener("click", swapImg1_2,false);

function swapImg1_3() {
    id = document.getElementById("img-1.3")
    src = id.getAttribute("src")
    if (src == "assets/1.3.svg") {
        console.log("RUN")
        id.setAttribute("src", "assets/1.3.1.svg")
    } else if (src == "assets/1.3.1.svg") {
        id.setAttribute("src", "assets/1.3.2.svg")
    } else if (src == "assets/1.3.2.svg") {
        id.setAttribute("src", "assets/1.3.svg")
    }
}
document.getElementById("img-1.3").addEventListener("click", swapImg1_3,false);