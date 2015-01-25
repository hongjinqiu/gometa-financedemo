#!/usr/bin/python
#encoding=utf8

import re,os

def getFieldStruct(text):
    result = {
        "text": text,
        "fieldNameValueDict": {},
    }
    result["id"] = re.findall(r'id="(.*?)"', text)[0]
    fieldLi = re.findall(r'<(.*?)>(.*?)</.*?>', text)
    for item in fieldLi:
        result["fieldNameValueDict"][item[0]] = item[1]
    
    return result

def getBusinessField():
    fIn = open("business_fieldpool.xml", 'r')
    content = fIn.read()
    fIn.close()
    
    fieldLi = re.findall(r'(?msi)<field .*?>.*?</field>', content)
    result = {}
    for item in fieldLi:
        fieldDict = getFieldStruct(item)
        result[fieldDict["id"]] = fieldDict
    return result

def deal():
    fIn = os.popen("find . -name 'ds_*.xml'")
    content = fIn.read()
    fIn.close()
    
    li = content.split("\n")
    li2 = [item for item in li if item]
    li = li2
    
    businessField = getBusinessField()
    
    for itemPath in li:
        fIn = open(itemPath, 'r')
        content = fIn.read()
        fIn.close()
        
        fieldLi = re.findall(r'(?msi)<field .*?>.*?</field>', content)
        for item in fieldLi:
            fieldDict = getFieldStruct(item)
            if businessField.get(fieldDict["id"]):
                tmpBusinessField = businessField[fieldDict["id"]]["fieldNameValueDict"]
                for k,v in fieldDict["fieldNameValueDict"].items():
                    if tmpBusinessField.get(k) and tmpBusinessField[k] == v:
                        fieldDict["text"] = fieldDict["text"].replace("<" + k + ">" + v + "</" + k + ">", "<" + k + ">the same</" + k + ">")
            
            content = content.replace(item, fieldDict["text"])
            
        fOut = open(itemPath, 'w')
        fOut.write(content)
        fOut.close()
        
deal()

