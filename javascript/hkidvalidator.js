function checkLength(hkid) {
  var len = hkid.length
  var hkid = hkid.toUpperCase()
  if (len === 8) {
    return " " + hkid
  } else if (len === 9) {
    return hkid
  } else {
    return ""
  }
}

function hkidSum(hkid) {
  var sum = 0
  for (i = 0; i < 8; i++) {
    if (hkid[i] === ' ') {
      sum += 36 * 9
    } else if (hkid[i].search(/[^A-Z]/)) {
      sum += (hkid.charCodeAt(i) - 'A'.charCodeAt(0) + 10) * (9 - i)
    } else {
      sum += (hkid.charCodeAt(i) - '0'.charCodeAt(0)) * (9 - i)
    }
  }
  return sum
}

function checkDigit(hkid) {
  if (hkid[8] === 'A') {
    return 10
  } else {
    return Number(hkid[8])
  }
}

function checkHKID(hkid) {
  s = checkLength(hkid)
  if (s != "") {
    sum = hkidSum(s)
    mod = sum % 11
    cd = checkDigit(s)
    if (mod === 0) {
      return cd === 0
    } else {
      return cd === 11 - mod
    }
  } else {
    return false
  }
}

module.exports = { checkHKID }
