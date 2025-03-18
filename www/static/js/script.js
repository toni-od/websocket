window.onload = init;

const state = {
    root: undefined,
}

function init() {
    root = document.querySelector("#root");
    root.innerHTML = "Websocket";
}
