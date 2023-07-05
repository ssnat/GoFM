export default function GetWebsocketUrl() {
    if(!window) return '';
    const protocol = window.location.protocol === "https:" ? "wss://" : "ws://";
    const hostname = window.location.hostname;
    const port = window.location.port || (protocol === "wss://" ? "443" : "80");
    return protocol + hostname + ":" + port + "/ws";
}