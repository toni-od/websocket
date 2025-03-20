window.onload = init;

const state = {
    root,
}

function init() {
    state.root = document.querySelector("#root");
    state.root.innerHTML = "Websocket";

    if(!window["WebSocket"]) {
        console.log("Browser does not support Websocket");
    } else {
        new WebSocket(`ws://${window.location.host}/ws`);
    }
}
