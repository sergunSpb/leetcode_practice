def findValidDiscountCoupons(discounts):
    # Write your code here
    hist_val = dict()
    hist_inval = dict()
    
    def helper(c):
        if len(c)  == 0:
            return True
        if len(c) %2 == 1:
            return False
        if c in hist_val:
            return True
        if c in hist_inval:
            return False
        if len(c) == 2:
            return c[0] == c[1]
        p = 2
        while p <len(c):
            r1 = helper(c[0:p])
            r2 = helper(c[p:])
            if r1 :
                hist_val[c[0:p]] = None
            else:
                hist_inval[c[0:p]] = None
            if r2 :
                hist_val[c[p:]] = None
            else:
                hist_inval[c[p:]] = None
            if r1 and r2:
                hist_val[c] = None
                return True        
            p +=2
        
        if c[0] == c[len(c)-1]:
            if helper(c[1:len(c)-1]):
                hist_val[c] = True
                return True

        hist_inval[c] = None
        return False
    res = []    
    for i in discounts:
        res.append(helper(i))
    return [ 1 if x else 0 for x in res ]  

print(findValidDiscountCoupons(["aabb","aa","abba","ffaawwxxx",""]))