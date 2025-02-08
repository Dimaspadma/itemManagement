import dayjs from 'dayjs'
import Response from "./response"
import Long from "long"
import type { Timestamp } from '../google/protobuf/timestamp';

export function create(username: string): (Uint8Array | null) {
  const date = new Date()
  const seconds = Math.floor(date.getTime() / 1000); // detik sejak Unix Epoch
  const nanos = (date.getMilliseconds() * 1000000); // konversi ms ke nanos
  
  const user: Response.awesomeProject.User = new Response.awesomeProject.User({
    id: -1,
    username,
    createdAt: {
      seconds,
      nanos,
    }
  })
  const err = Response.awesomeProject.User.verify(user)
  if (err){
    console.error(err)
    return null
  }
  const message = Response.awesomeProject.User.create(user)
  const buffer = Response.awesomeProject.User.encode(message).finish()
  console.log(buffer)
  return buffer
}

export function decodeResponse(data: any) : Response.awesomeProject.Response {
  const response: Response.awesomeProject.Response = Response.awesomeProject.Response.decode(new Uint8Array(data))
  return response
}

export function formatTimestamp(timeStamp: Response.google.protobuf.ITimestamp | null | undefined): Date {
  let seconds: number = toNumber(timeStamp?.seconds)
  return dayjs.unix(seconds).toDate()
}

export function formatTimestamp2(timestamp: Timestamp | undefined): Date {
  let seconds: number = Number(timestamp?.seconds)
  return dayjs.unix(seconds).toDate()
}

function toNumber(value: any) {
  if (Long.isLong(value)) {
    return value.toNumber(); // Konversi Long ke number
  }
  return value; // Kalau udah number, langsung return
}