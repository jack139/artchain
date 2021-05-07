# coding:utf-8
import sys
import urllib3, json, base64, time, hashlib
from datetime import datetime

urllib3.disable_warnings()

#with open("doc/exchainge.png", 'rb') as f:
#with open("2021030117343322954.zip", 'rb') as f:
#    img_data = f.read()
#img_data = base64.b64encode(img_data).decode('utf-8')

# 生成参数字符串
def gen_param_str(param1):
    param = param1.copy()
    name_list = sorted(param.keys())
    if 'data' in name_list: # data 按 key 排序, 中文不进行性转义，与go保持一致
        param['data'] = json.dumps(param['data'], sort_keys=True, ensure_ascii=False, separators=(',', ':'))
    return '&'.join(['%s=%s'%(str(i), str(param[i])) for i in name_list if str(param[i])!=''])


if __name__ == '__main__':
    if len(sys.argv)<2:
        print("usage: python3 %s <host> <port>" % sys.argv[0])
        sys.exit(2)

    hostname = sys.argv[1]
    port = sys.argv[2]

    body = {
        'version'  : '1',
        'sign_type' : 'SHA256', 
        'data'     : {
            #'chain_addr'   : 'bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004', # test1
            #'chain_addr'   : 'bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7', # test1

            'login_name' : 'test3',
            #'user_type' : 'TRD',
            #'email' : '111112@qq.com',
            #'bank_acc_name' : '1test bank',
            #'referrer': 'bid1art111111111',
            'chain_addr'   : 'bid1art16zs5zpmsw5wezyrpnls76ytdy7ws2zpqan9ey9', # test3
            'mystery' : 'denial move indoor quick monkey share split fog expose orbit merge flash twelve vicious salmon toast gold unusual have purchase time dune satoshi spoil',

            #'height' : '985',

            #'id' : '0',

            #'desc' : '测试物品3',
            #'date' : '1911s',
            #'base_price' : '$2001',
            #'owner_addr' : 'bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7',

            #'page' : 1,
            #'limit' : 10,

            #'item_id' : "3",
            #'reviewer_addr' : "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
            #'detail' : "aaaaabbbbbb 哈哈1111111",

            #'seller_addr' : "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
            #'auction_house_id' : "1",
            #'item_id' : '2',
            #'reserved_price' : '2000',

            #'buyer_addr' :  "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
            #'auction_id' : "1",
            #'item_id' : "2",
            #'trans_type' : "BID",
            #'hammer_time' : "2021-01-01",
            #'hammer_price' : "1000",
            #'details' : "测试测试测试"
        }
    }

    secret = 'MjdjNGQxNGU3NjA1OWI0MGVmODIyN2FkOTEwYTViNDQzYTNjNTIyNSAgLQo='
    appid = '4fcf3871f4a023712bec9ed44ee4b709'
    unixtime = int(time.time())
    body['timestamp'] = unixtime
    body['appid'] = appid

    param_str = gen_param_str(body)
    sign_str = '%s&key=%s' % (param_str, secret)

    if body['sign_type'] == 'SHA256':
        sha256 = hashlib.sha256(sign_str.encode('utf-8')).hexdigest().encode('utf-8')
        signature_str =  base64.b64encode(sha256).decode('utf-8')
    else: # SM2
        #signature_str = sm2.SM2withSM3_sign_base64(sign_str)
        pass

    #print(sign_str.encode('utf-8'))
    #print(sha256)
    #print(signature_str)

    body['sign_data'] = signature_str

    body = json.dumps(body)
    print(body)

    pool = urllib3.PoolManager(num_pools=2, timeout=180, retries=False)

    host = 'http://%s:%s'%(hostname, port)
    #url = host+'/api/test'
    #url = host+'/api/r1/biz/user/register'
    #url = host+'/api/r1/biz/user/modify'
    #url = host+'/api/r1/biz/item/new'
    #url = host+'/api/r1/biz/item/modify'
    #url = host+'/api/r1/biz/review/new'
    #url = host+'/api/r1/biz/review/modify'
    #url = host+'/api/r1/biz/auction/new'
    #url = host+'/api/r1/biz/auction/modify'
    #url = host+'/api/r1/biz/trans/new'

    #url = host+'/api/r1/query/block/rawdata'
    #url = host+'/api/r1/query/user/credit_balance'
    #url = host+'/api/r1/query/user/info'
    url = host+'/api/r1/query/user/verify'
    #url = host+'/api/r1/query/item/info'
    #url = host+'/api/r1/query/item/list'
    #url = host+'/api/r1/query/review/list'
    #url = host+'/api/r1/query/auction/info'
    #url = host+'/api/r1/query/auction/list'
    #url = host+'/api/r1/query/trans/info'
    #url = host+'/api/r1/query/trans/list'
    

    start_time = datetime.now()
    r = pool.urlopen('POST', url, body=body)
    print('[Time taken: {!s}]'.format(datetime.now() - start_time))

    print(r.status)
    if r.status==200:
        #print(json.loads(r.data.decode('utf-8')))
        print(r.data.decode('utf-8'))
    else:
        print(r.data)
