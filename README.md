gometa-financedemo是gometa框架的演示，其功能是实现一个简单的财务出纳系统。

演示项目url地址:http://112.126.81.24/


    li = []
    li.append({
        'name': '基础资料',
        'isLeaf': 2,
    })
    li.append({
        'name': '系统用户', 
        'url': '/console/listschema?@name=SysUser&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '项目类别', 
        'url': '/console/listschema?@name=ArticleType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '项目', 
        'url': '/console/listschema?@name=Article&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '供应商类别', 
        'url': '/console/listschema?@name=ProviderType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '供应商', 
        'url': '/console/listschema?@name=Provider&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '币别', 
        'url': '/console/listschema?@name=CurrencyType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '银行资料', 
        'url': '/console/listschema?@name=Bank&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '计量单位', 
        'url': '/console/listschema?@name=MeasureUnit&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '客户类别', 
        'url': '/console/listschema?@name=CustomerType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '客户', 
        'url': '/console/listschema?@name=Customer&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '税率类别', 
        'url': '/console/listschema?@name=TaxType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '结算方式', 
        'url': '/console/listschema?@name=BalanceType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '收入费用类别', 
        'url': '/console/listschema?@name=IncomeType&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '收入费用项目', 
        'url': '/console/listschema?@name=IncomeItem&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '现金账户', 
        'url': '/console/listschema?@name=CashAccount&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '银行账户', 
        'url': '/console/listschema?@name=BankAccount&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '会计期', 
        'url': '/console/listschema?@name=AccountingPeriod&cookie=false',
        'isLeaf': 1,
    })
    
    li.append({
        'name': '单据类型参数',
        'isLeaf': 2,
    })
    li.append({
        'name': '单据类型参数', 
        'url': '/console/listschema?@name=BillTypeParameter&cookie=false',
        'isLeaf': 1,
    })
    
    li.append({
        'name': '系统参数',
        'isLeaf': 2,
    })
    li.append({
        'name': '系统参数', 
        'url': '/console/listschema?@name=SystemParameter&cookie=false',
        'isLeaf': 1,
    })
    
    li.append({
        'name': '初始化',
        'isLeaf': 2,
    })
    li.append({
        'name': '现金账户初始化', 
        'url': '/console/formschema?@name=CashAccountInit&formStatus=view&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '银行账户初始化', 
        'url': '/console/formschema?@name=BankAccountInit&formStatus=view&cookie=false',
        'isLeaf': 1,
    })
    
    li.append({
        'name': '单据',
        'isLeaf': 2,
    })
    li.append({
        'name': '收款单', 
        'url': '/console/listschema?@name=GatheringBill&cookie=false',
        'isLeaf': 1,
    })
    li.append({
        'name': '付款单', 
        'url': '/console/listschema?@name=PayBill&cookie=false',
        'isLeaf': 1,
    })
    
    li.append({
        'name': '报表',
        'isLeaf': 2,
    })
    li.append({
        'name': '资金汇总表', 
        'url': '/console/formschema?@name=AccountInOut',
        'isLeaf': 1,
    })
