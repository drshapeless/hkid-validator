def checkLength(hkid):
    hkid = hkid.upper()
    if len(hkid) == 8 \
       and hkid[0].isalpha() \
       and hkid[1:7].isnumeric() \
       and (hkid[7] == 'A' or hkid[7].isnumeric()):
        return ' ' + hkid

    elif len(hkid) == 9 \
            and hkid[0:2].isalpha() \
            and hkid[2:8].isnumeric() \
            and (hkid[8] == 'A' or hkid[8].isnumeric()):
        return hkid
    else:
        return ''


def hkidSum(hkid):
    s = 0
    for i in range(8):
        if hkid[i] == ' ':
            s += 36 * 9
        elif hkid[i].isalpha():
            s += (ord(hkid[i]) - ord('A') + 10) * (9 - i)
        else:
            s += int(hkid[i]) * (9 - i)
    return s


def checkDigit(hkid):
    if hkid[8] == 'A':
        return 10
    else:
        return int(hkid[8])


def checkHKID(hkid):
    s = checkLength(hkid)
    if s == '':
        return False
    else:
        mod = hkidSum(s) % 11
        checkdigit = checkDigit(s)
        if checkdigit == 0:
            return mod == checkdigit
        else:
            return checkdigit == 11 - mod
