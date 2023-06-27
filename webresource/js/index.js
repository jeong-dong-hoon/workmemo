'use strict'

let topmenu = document.getElementById("topmenu");
let exitbtn = document.querySelector(".exitbtn");
let plusbtn = document.querySelector(".plusbtn");
let textboard = document.getElementById("textboard");

document.addEventListener('mouseover',(evt)=>{
    topmenu.classList.toggle('on')
    topmenu.style.top = '0px'
})
document.addEventListener('mouseout',(evt)=>{
    topmenu.classList.remove('on')
    topmenu.style.top = '-20px'
})