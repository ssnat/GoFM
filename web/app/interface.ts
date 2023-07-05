export type IFMInfo = {
  title: string;
  artist: string;
  cover: string;
  url: string;
  sampleRate: string;
  bitRate: string;
}

export type IServerInfo = {
  name: string;
  version: string;
  time: number;
}


export type IWebsocketStatus = 'connecting' | 'disconnected' | 'connected' | 'error';