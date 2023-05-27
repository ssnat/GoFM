"use client";

import Link from 'next/link'
import {useEffect, useState} from "react";
export default function Home() {

  const [musicInfo, setMusicInfo] = useState({
    init: false,
    title: '',
    artist: '',
    cover: '',
  });


  useEffect(() => {
    fetch('http://127.0.0.1:8090/api/fm/info')
      .then(res => res.json())
      .then(res => {
        const {title, artist, cover} = res.data;
        setMusicInfo({
          init: true,
          title,
          artist,
          cover,
        })
      })
      .catch(console.error)
  }, [])

  return (
    <main className="flex justify-center items-center h-screen w-screen bg-gray-100">
      <div className="w-96 mx-auto max-w-full p-4">
        <div className="mb-2">GoFM</div>
        <div className="bg-white rounded-md overflow-hidden h-24 flex mb-2">
          <div className="w-24 h-24 flex-shrink-0 ">
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img src="http://127.0.0.1:8090/api/fm/info/cover" className="h-full w-full" alt="12"/>
          </div>
          <div className="pt-3 pb-3 pl-3 pr-3 flex-1 overflow-hidden">
            <div className="h-6 truncate w-full font-bold text-md text-gray-700 mb-1">
              {musicInfo.title}
            </div>
            <div className="h-5 truncate w-full text-sm mb-1 text-gray-700">
              {musicInfo.artist}
            </div>
            <div className="h-6 truncate w-full text-sm text-gray-700">
              00:00:01
            </div>
          </div>
        </div>
        <div className="text-center text-xs text-gray-600">
          <Link href={"https://github.com/pxgo/GoFM"} target={"_blank"}>GoFM v0.5.1</Link>
        </div>
      </div>
    </main>
  )
}
