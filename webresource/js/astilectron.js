'use strict'

let message = {
    Code : ""
    ,Content : []
}
document.addEventListener('astilectron-ready', function() {

    astilectron.onMessage(function(message) {

        if (message === "hello") {
            return "world";
        }
    });
})
document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        console.log(message)
    });
    exitbtn.addEventListener('click',()=>{
        message.Code = "done";
        astilectron.sendMessage(message, ()=>{});
    })
    
    plusbtn.addEventListener('click',()=>{
        message.Code = "save";
        astilectron.sendMessage(message,()=>{});
    }

    )
   
})