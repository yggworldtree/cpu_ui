/**   
 * api接口统一管理
 */
 import { get, post } from './http'

 export const cpuProcs = p => get('cpu-procs', p);
 export const cpuMem = p => get('cpu-mem', p);