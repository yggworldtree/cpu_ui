// 在http.js中引入axios
import axios from 'axios'; // 引入axios
// 环境的切换
const request = axios.create({
  baseURL: '/',   // 基础路径，默认是/ ，如果改了，会自动添加到你请求url前面
  timeout: 10000   // 请求超时，5000毫秒
})

request.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';


/**
 * get方法，对应get请求
 * @param {String} url [请求的url地址]
 * @param {Object} params [请求时携带的参数]
 */
 export function get(url, params){    
  return new Promise((resolve, reject) =>{        
    request.get(url, {            
          params: params        
      }).then(res => {
          resolve(res.data);
      }).catch(err =>{
          reject(err.data)        
  })    
});}

/** 
 * post方法，对应post请求 
 * @param {String} url [请求的url地址] 
 * @param {Object} params [请求时携带的参数] 
 */
 export function post(url, params) {
  return new Promise((resolve, reject) => {
    request.post(url, params)
      .then(res => {
          resolve(res.data);
      })
      .catch(err =>{
          reject(err.data)
      })
  });
}