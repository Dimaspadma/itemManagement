import paketData from "../assets/dummy/paketData.json"

export interface Paket {
  name: string,
  status: string,
  photo: string
}

export const fetchPaketData = () => {
  return new Promise<Paket[]>( (resolve) => {
    setTimeout( () => {
      resolve(paketData);
    }, 500)
  })
}