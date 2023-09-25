package service

import (
	"context"
	"math/rand"
	pb "verify-code/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}

func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randCode("0123456789", l, 4)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyz", l, 5)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyz", l, 6)
	default:
	}
	return ""
}

// 简单实现
//func randCode(chars string, l int) string {
//	charsLen := len(chars)
//	result := make([]byte, l)
//	for i := 0; i < l; i++ {
//		randIndex := rand.Intn(charsLen)
//		result[i] = chars[randIndex]
//	}
//	return string(result)
//}

// 优化实现
// 一次随机多个随机位，分部分多次使用
func randCode(chars string, l, idxBits int) string {
	// 计算有效的二进制数位，基于 chars 的长度
	// 推荐写死，因为 chars 固定，对应位数长度也固定
	// idxBits = len(fmt.Sprintf("%b", len(chars)))

	// 形成掩码
	// 例如使用低六位 00000000000111111
	idxMask := 1<<idxBits - 1 // 00001000000 - 1 = 00000111111
	// 63 位可用多少次
	idxMax := 63 / idxBits
	// 结果
	result := make([]byte, l)
	// 生成随机字符
	// cache 随机位缓存
	// remain 当前还可以用几次
	for i, cache, remain := 0, rand.Int63(), idxMax; i < l; {
		// 如果使用的剩余次数为 0 ，重新随机获取
		if 0 == remain {
			cache, remain = rand.Int63(), idxMax
		}
		// 利用掩码获取有效部位的随机数位
		if randIndex := int(cache & int64(idxMask)); randIndex < len(chars) {
			result[i] = chars[randIndex]
			i++
		}
		// 使用下一组随机位
		cache >>= idxBits
		// 减少一次使用次数
		remain--
	}
	return string(result)
}
