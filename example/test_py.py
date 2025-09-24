import pyctp

cq = pyctp.Quote("./log_quote222/")
cq.CreateApi()
cq.CreateSpi()
print(cq.GetApiVersion())

ct = pyctp.Trade("./log_trade222/")
ct.CreateApi()
ct.CreateSpi()
print(ct.GetApiVersion())

