/**
 * Copyright (c) yangzhao 2022/3/25
 *
 * File name: functions.go
 * Created at: 2022/3/25 9:49 AM
 * Author: yangzhao yzha5@163.com
 *
 * Description: Define status functions
 */

package status

// merge maps for type map[string]string
func merge(ms ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
