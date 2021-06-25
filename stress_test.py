# -*- coding: utf-8 -*-
#
import os
import time, hashlib, json, sys, random, base64
import threading
import urllib3

urllib3.disable_warnings()


TEST_SERVER = [
    '127.0.0.1:8888'
]


BODY = {
    'version'  : '1',
    'sign_type' : 'SHA256', 
    'data'     : {
        'caller_addr' : 'bid1art17qppfv5k29r9txqu8sj3l6vfwtt90rr82r9gt7',
        'id' : '3',
        'detail' : '压力测试',
    }
}


SCORE = {}

# 生成参数字符串
def gen_param_str(param1):
    param = param1.copy()
    name_list = sorted(param.keys())
    if 'data' in name_list: # data 按 key 排序, 中文不进行性转义，与go保持一致
        param['data'] = json.dumps(param['data'], sort_keys=True, ensure_ascii=False, separators=(',', ':'))
    return '&'.join(['%s=%s'%(str(i), str(param[i])) for i in name_list if str(param[i])!=''])


def main_loop(tname, test_server):
    global SCORE

    body = BODY.copy()

    secret = 'MjdjNGQxNGU3NjA1OWI0MGVmODIyN2FkOTEwYTViNDQzYTNjNTIyNSAgLQo='
    appid = '4fcf3871f4a023712bec9ed44ee4b709'
    unixtime = int(time.time())
    body['timestamp'] = unixtime
    body['appid'] = appid

    param_str = gen_param_str(body)
    sign_str = '%s&key=%s' % (param_str, secret)

    sha256 = hashlib.sha256(sign_str.encode('utf-8')).hexdigest().encode('utf-8')
    signature_str =  base64.b64encode(sha256).decode('utf-8')

    #print(sign_str)

    body['sign_data'] = signature_str

    body = json.dumps(body)

    #print(body)

    pool = urllib3.PoolManager(num_pools=100, timeout=180, retries=False)
    host = 'http://%s'%(test_server)
    url = host+'/api/r1/biz/item/modify'

    tick_start = time.time()

    try:
        print(url)
        #print header

        r = pool.urlopen('POST', url, body=body)
        #print(r.data)
        #print(r.status)

        tick_end = time.time()

        time_used =  int((tick_end-tick_start)*1000)
        SCORE[tname] += time_used

        if r.status!=200:
            print(tname, test_server, '!!!!!! HTTP ret=', r.status, 'time_used=', time_used)
        else:
            r2 = json.loads(r.data.decode('utf-8'))
            print(tname, test_server, '200', 'time_used=', time_used, 
                'in_data_len=', len(body), 'out_data_len=', len(r.data), 'code=', r2['code'])
    except Exception as e:
        print("异常: %s : %s" % (e.__class__.__name__, e))

class MainLoop(threading.Thread):
    def __init__(self, rounds):
        threading.Thread.__init__(self)
        self._tname = None
        self._round = rounds

    def run(self):
        global count, mutex, SCORE
        self._tname = threading.currentThread().getName()
        SCORE[self._tname] = 0        

        print('Thread - %s started.' % self._tname)

        #while 1:
        for x in range(0, self._round):
            for y in TEST_SERVER:
                main_loop(self._tname, y)

            # 周期性打印日志
            time.sleep(random.randint(0,1))
            sys.stdout.flush()


if __name__=='__main__':
    if len(sys.argv)<3:
        print("usage: python stress_test.py <thread_num> <round_per_thread>")
        sys.exit(2)

    print("STRESS TEST started: " , time.ctime())

    thread_num = int(sys.argv[1])
    round_per_thread = int(sys.argv[2])

    #线程池
    threads = []
        
    # 创建线程对象
    for x in range(0, thread_num):
        threads.append(MainLoop(round_per_thread))
    
    # 启动线程
    for t in threads:
        t.start()

    # 等待子线程结束
    for t in threads:
        t.join()  

    total = 0
    for i in SCORE.keys():
        total += SCORE[i]
        print('%s - %.3f'%( i, SCORE[i]/round_per_thread ))

    print('Average: %.3f'%(total/(thread_num*round_per_thread)) )

    print("STRESS TEST exited: ", time.ctime())
