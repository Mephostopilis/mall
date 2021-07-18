import axios from 'axios'
import qs from 'qs'

export function upload(data) {
  return axios({
    method: 'post',
    url: "http://127.0.0.1:8080/",
    data: data
  })
}

export function download(data) {
  return axios({
    method: 'get',
    url: 'http://bit.ly/2mTM3nY',
    responseType: 'stream'
  })
}
