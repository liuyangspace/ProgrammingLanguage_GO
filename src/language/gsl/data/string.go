package data

import (
	"strings"
	"strconv"
	"regexp"
)

/*
字符及字符串
	界定符:
		""	// (双引号)解释转义
		``	// (反引号)不解释转义，\n等会被原样输出。
	字符串拼接符: +
	字节转换:
		s := "hello"
		c := []byte(s)	// 字符串->字节数组
		s2 := string(c) // 字节数组->字符串
	正则:
		参考 regexp 包

 */
// 字符索引
func HasPrefix(s, prefix string) bool { return strings.HasPrefix(s, prefix) }//判断字符串 s 是否以 prefix 开头
func HasSuffix(s, suffix string) bool { return strings.HasSuffix(s, suffix) }//判断字符串 s 是否以 suffix 结尾
func Contains(s, substr string) bool { return strings.Contains(s, substr) }//判断字符串 s 是否包含 substr
func Index(s, str string) int { return strings.Index(s, str) }//返回字符串 str 在字符串 s 中的索引，-1 表示字符串 s 不包含字符串 str
func LastIndex(s, str string) int { return strings.LastIndex(s, str) }//返回字符串 str 在字符串 s 中最后出现位置的索引，-1 表示字符串 s 不包含字符串 str
func IndexRune(s string, r rune) int { return strings.IndexRune(s, r) }//查询非 ASCII 编码的字符在父字符串中的位置
// 替换
func Replace(str string, old string, new string, n int) string { return strings.Replace(str, old, new, n) }//将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new
func TrimSpace(s string) string { return strings.TrimSpace(s) }//剔除字符串开头和结尾的空白符号
func Trim(s string, del string) string { return strings.Trim(s,del) }//剔除指定字符
func TrimLeft(s string, cutset string) string { return strings.TrimLeft(s,cutset) }//剔除开头
func TrimRight(s string, cutset string) string { return strings.TrimRight(s,cutset) }//剔除结尾
// 统计
func Count(s, str string) int { return strings.Count(s, str) }//计算字符串 str 在字符串 s 中出现的非重叠次数
// 切割 拼接
func Fields(s string) []string { return strings.Fields(s) }//利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块
func Split(s, sep string) []string { return strings.Split(s,sep) }//用于自定义分割符号来对指定字符串进行分割
func Join(sl []string, sep string) string { return strings.Join(sl,sep) }//将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
// 转换
func Repeat(s string, count int) string { return strings.Repeat(s, count) }//重复 count 次字符串 s 并返回一个新的字符串
func ToLower(s string) string { return strings.ToLower(s) }//将字符串中的 Unicode 字符全部转换为相应的小写字符
func ToUpper(s string) string { return strings.ToUpper(s) }//将字符串中的 Unicode 字符全部转换为相应的大写字符
func Itoa(i int) string { return strconv.Itoa(i) }//返回数字 i 所表示的字符串类型的十进制数
func FormatInt(i int64, base int) string { return strconv.FormatInt(i,base) }//
func FormatFloat(f float64, fmt byte, prec, bitSize int) string { return strconv.FormatFloat(f,fmt,prec,bitSize) }//将 64 位浮点型的数字转换为字符串
func Atoi(s string) (int, error) { return strconv.Atoi(s) } //将字符串转换为 int 型
func ParseFloat(s string, bitSize int) (float64, error) { return strconv.ParseFloat(s,bitSize) } //将字符串转换为 float64 型
// 正则
func MatchString(pattern string, s string) (matched bool, err error){ return regexp.MatchString(pattern,s) } //
// 类型
func NewReader(s string) *strings.Reader { return strings.NewReader(s) }