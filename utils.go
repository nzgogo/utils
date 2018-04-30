package utils

import "strings"

func MapKeyToCamelCase(m map[string]interface{}) map[string]interface{} {
  for k, v := range m {
    l := strings.ToLower(k)
    s := strings.Split(l, "_")
    nk := ""
    for i, kp := range s {
      if i == 0 {
        nk = nk + kp
      } else {
        nk = nk + strings.Title(kp)
      }
    }
    m[nk] = m[k]
    delete(m, k)
  }
  return m
}
