window.onload = init;

const state = {
    root,
}

function init() {
    state.root = document.querySelector("#root");
    state.root.innerHTML = "Websocket";
}
