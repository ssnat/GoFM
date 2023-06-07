"use client";
import Head from 'next/head'
import Link from 'next/link'
import React, {useCallback, useEffect, useState} from "react";
import {MdPauseCircle, MdPlayCircle} from 'react-icons/md'

type IPlayerStatus = "loading"|'playing'|'stopped';

/*export async function generateMetadata({ params }: Props): Promise<Metadata> {
  // read route params then fetch data

  // return an object
  return {
    title: "AAAA",
    description: "BBB",
  };
}*/

class Home extends React.Component<any, {
  errorMessage: string;
  runningTime: string;
  timer: null | NodeJS.Timeout;
  playerStatus: IPlayerStatus;
  serverInfo:null | {
    name: string;
    version: string;
    time: number;
  };
  fmInfo: null | {
    title: string;
    artist: string;
    cover: string;
    url: string;
    sampleRate: string;
    bitRate: string;
  };
}>{
  audioRef = React.createRef<HTMLAudioElement>();
  constructor(props: any) {
    super(props)
    this.state = {
      errorMessage: '',
      runningTime: '00:00:00',
      timer: null,
      serverInfo: null,
      fmInfo: null,
      playerStatus: 'stopped',
    }
  }

  componentDidMount() {
    this.setTimerToUpdateRunningTime();
    this.initInfo().catch(err => {
      this.setState({
        errorMessage: err.message || err.toString()
      })
    });
  }

  componentWillUnmount() {
    this.clearTimer();
  }

  setPlayerStatus() {
    if(this.state.playerStatus === 'playing') {
      this.stop();
    } else if(this.state.playerStatus === 'stopped') {
      this.play();
    } else if(this.state.playerStatus === 'loading') {
      this.stop();
    }
  }

  play() {
    if(!this.audioRef.current) return
    this.audioRef.current.play()
      .then(() => {
        this.setState({
          playerStatus: 'playing'
        })
      })
      .catch(err => {
        this.setState({
          playerStatus: 'stopped'
        })
      })
  }

  stop() {
    if(!this.audioRef.current) {
      return;
    }
    this.audioRef.current.pause();
    this.setState({
      playerStatus: 'stopped'
    });
  }

  initInfo() {
    return fetch('/api/info')
      .then(res => res.json())
      .then(res => {
        const {name, version, time, FMInfo} = res.data;
        const {cover, title, artist, sampleRate, bitRate, url} = FMInfo;
        this.setState({
          serverInfo: {
            name,
            version,
            time
          },
          fmInfo: {
            cover: cover + '?t=' + Date.now(),
            title,
            artist,
            sampleRate,
            bitRate,
            url,
          }
        })
      });
  }

  setTimerToUpdateRunningTime() {
    const timer = setTimeout(() => {
      if(this.state.serverInfo) {
        this.setState({
          runningTime: formatTime(Date.now() - this.state.serverInfo.time)
        })
      }
      this.setTimerToUpdateRunningTime();
    }, 1000);
  }

  clearTimer() {
    if(this.state.timer) {
      clearTimeout(this.state.timer)
    }
  }

  render() {
    const {
      runningTime,
      errorMessage,
      serverInfo,
      fmInfo,
      playerStatus
    } = this.state;
    return (
      <main className="flex justify-center items-center h-screen w-screen bg-gray-100">
        <Head>
          <title>{serverInfo? serverInfo.name: 'GoFM'}</title>
        </Head>
        {
          (() => {
            if(errorMessage) {
              return <TextInfo>{errorMessage}</TextInfo>
            } else if(!serverInfo || !fmInfo) {
              return <TextInfo>Loading...</TextInfo>
            } else {
              return (
                <>
                  <div style={{
                    backgroundImage: `url(${fmInfo.cover})`,
                    filter: `blur(100px)`,
                    WebkitFilter: `blur(100px)`,
                  }} className={"h-full w-full fixed left-0 top-0 z-10 bg-center bg-cover"}>
                    <div style={{
                      height: '100%',
                      width: '100%',
                      backgroundColor: 'rgba(255,255,255,0.6)',
                    }}></div>
                  </div>
                  <div className="w-96 mx-auto max-w-full p-4 z-20">
                    <div className="mb-2 flex justify-between items-center">
                      <div className={"text-left"}>{serverInfo.name}</div>
                      <div style={{
                        fontFamily: 'Arial, sans-serif'
                      }} className={"text-right text-gray-700 text-sm"}>{runningTime}</div>
                    </div>
                    <div className="bg-white rounded-md overflow-hidden h-24 flex mb-2 shadow-sm">
                      <div className="w-24 h-24 flex-shrink-0 ">
                        {/* eslint-disable-next-line @next/next/no-img-element */}
                        <img src={fmInfo.cover} className="h-full w-full" alt="12"/>
                      </div>
                      <div className="pt-3 pb-3 pl-3 pr-3 flex-1 overflow-hidden">
                        <div className="h-6 truncate w-full font-bold text-md text-gray-700 mb-0">
                          {fmInfo.title}
                        </div>
                        <div className="h-5 truncate w-full text-sm mb-1 text-gray-700">
                          {fmInfo.artist}
                        </div>
                        <div className="h-6 truncate w-full text-sm text-gray-700 flex">
                          <div className={"h-6 w-6 flex justify-center items-center cursor-pointer select-none"} onClick={() => {
                            this.setPlayerStatus();
                          }}>
                            {
                              playerStatus === 'stopped'? (
                                <MdPlayCircle size={30} />
                              ):(
                                <MdPauseCircle size={30} />
                              )
                            }
                          </div>
                        </div>
                      </div>
                    </div>
                    <audio src={fmInfo?.url} ref={this.audioRef}></audio>
                    <div className="text-center text-xs text-gray-600">
                      <Link href={"https://github.com/pxgo/GoFM"} target={"_blank"}>GoFM {serverInfo.version}</Link>
                    </div>
                  </div>
                </>
              )
            }
          })()
        }
      </main>
    )
  }
}

export default Home;

function TextInfo(props: {
  children: React.ReactNode
}) {
  return <div>{props.children}</div>
}

function formatTime(milliseconds: number) {
  let totalSeconds = Math.floor(milliseconds / 1000);
  let hours = Math.floor(totalSeconds / 3600);
  totalSeconds -= hours * 3600;
  let minutes = Math.floor(totalSeconds / 60);
  let seconds = totalSeconds - minutes * 60;
  return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
}