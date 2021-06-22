/**   
 * api接口统一管理
 */
 import { get, post } from './http'

 export const cpuProcs = p => get('procs', p);
 export const cpuMem = p => get('cpu-mem', p);
 export const queryWarns = p => get('warns', p);
 export const queryWarnln = p => get('warnln', p);
 export const queryWinfoln = p => get('winfoln', p);
 export const queryWinfos = p => get('winfos', p);