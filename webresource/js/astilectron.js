'use strict'

let message = {
    Code : ""
    ,Content : ""
}
document.addEventListener('astilectron-ready', function() {

    astilectron.onMessage(function(message) {
        console.log(message)
        if(message.Code == 'first') {
            textboard.contentWindow.document.querySelector('.content').innerHTML = ``;
        }
    });
})
document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        console.log(message)
    });
    exitbtn.addEventListener('click',()=>{
        message.Code = "done";
        message.Content = textboard.contentWindow.document.querySelector('.content').innerHTML;
        astilectron.sendMessage(message, ()=>{});
    })
    
    plusbtn.addEventListener('click',()=>{
        message.Code = "save";
        astilectron.sendMessage(message,()=>{});
    }

    )
   
})