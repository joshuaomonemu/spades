var socket = new WebSocket("ws://localhost:2020/ws")

function active() {
    document.getElementById('status').innerHTML = "Online";
}

console.log("Attempting connection");

socket.onopen = () => {
    console.log("Successfully connected");
    document.getElementById('status').innerHTML = "Online";
}
socket.onclose = (event) => {
    console.log("Socket closed: ", event)
    document.getElementById('status').innerHTML = "Offline";
}
socket.onmessage = (msg) => {
    console.log(msg);
}
socket.onerror = (error) => {
    console.log("Socket error:", error);
}
