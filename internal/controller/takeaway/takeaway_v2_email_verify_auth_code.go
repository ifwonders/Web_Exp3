package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV2) EmailVerifyAuthCode(ctx context.Context, req *v2.EmailVerifyAuthCodeReq) (res *v2.EmailVerifyAuthCodeRes, err error) {
	var authcodeRecord *entity.Authcode
	err = dao.Authcode.Ctx(ctx).Where(do.Authcode{Email: req.Email}).Scan(&authcodeRecord)

	if err != nil {
		return nil, err
	}
	var verifyResult = authcodeRecord.Authcode == req.Authcode && authcodeRecord.ExpiredTime.After(gtime.Now())

	res = &v2.EmailVerifyAuthCodeRes{
		VerifyResult: verifyResult,
	}
	return
}
