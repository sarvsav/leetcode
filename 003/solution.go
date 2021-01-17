// Problem: https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
    result := strings.Replace(address, ".", "[.]", -1)
    return result
}
