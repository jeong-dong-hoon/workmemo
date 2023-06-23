'use strict'

let message = {
    Code :'done'
}
document.addEventListener('astilectron-ready', function() {
    // This will listen to messages sent by GO
    astilectron.onMessage(function(message) {
        // Process message
        if (message === "hello") {
            return "world";
        }
    });
})
document.addEventListener('astilectron-ready', function() {
    exitbtn.addEventListener('click',()=>{
        astilectron.sendMessage(message, function(message) {
           
        });
    })
   
})