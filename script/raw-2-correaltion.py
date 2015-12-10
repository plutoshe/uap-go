import csv
class Dog:
    def __init__(self, id_1, actiontype, appid, detail, time):
        self.time = time
        self.id_1 = id_1
        self.appid = appid
        self.time = time
        self.detail = detail

def getKey(item):
    return item['time']

with open('text.csv', 'rb') as csvfile:
    spamreader = csv.reader(csvfile, delimiter=',', quotechar='"')
    # print string(spamreader)
    i =0

    ll = [{"ip":"61.149.193.238", "action":"open","time":1231312321}]
    # ll = []
    for row in spamreader:
        if row[5] =="":
            continue
        i = i + 1
        if i == 1:
            continue
        # if (i > 100):
        #     break
        # print row

        s = {'appid':row[2], "id":row[0],"action":row[1], 'detail':row[3],'time':int(row[4]), "ip":row[5], "ua": row[6]}
        # w = Dog(row[0], row[1], row[2], row[3], int(row[4]))
        ll.append(s)
    ll = sorted(ll, key=lambda x : x['time'])
    dd = {}
    with open('ans.csv', 'wb') as csvfile:
        spamwriter = csv.writer(csvfile, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
        for item in ll:
            # print item['time']
            if item["action"] == "userlink" or item["action"] == "userlink_wechat":
                dd[item['ip']] = item
            else:
                if (item["action"] == "open" or item["action"] == "install") and (item['ip'] in dd):
                    print item['ip']
                    # print dd
                    w1=[]
                    w2=[]
                    if (item['time']-dd[item['ip']]['time'] > 9000):
                        continue
                    for k,v in item.iteritems():
                        w1.append(v)
                    w1.append(item['time']-dd[item['ip']]['time'])
                    for k,v in dd[item['ip']].iteritems():
                        w2.append(v)
                    spamwriter.writerow(w1)
                    spamwriter.writerow(w2)
                    spamwriter.writerow([])

        # print dd




        # print '$$@#'.join(row)