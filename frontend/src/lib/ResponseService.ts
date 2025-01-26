import dayjs from 'dayjs'
import Response from "./response"
import Long from "long"

export function decodeResponse(data: any) : Response.awesomeProject.Response {
  const response: Response.awesomeProject.Response = Response.awesomeProject.Response.decode(new Uint8Array(data))
  return response
}

export function formatTimestamp(timeStamp: Response.google.protobuf.ITimestamp | null | undefined): Date {
  let seconds: number = toNumber(timeStamp?.seconds)
  return dayjs.unix(seconds).toDate()
}

function toNumber(value: any) {
  if (Long.isLong(value)) {
    return value.toNumber(); // Konversi Long ke number
  }
  return value; // Kalau udah number, langsung return
}