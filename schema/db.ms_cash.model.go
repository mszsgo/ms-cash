package schema

import (
	"time"
)

// 现金账户信息：ms_cash_account
type CashAccount struct {
	AccountNo string        `bson:"accountNo"` // 账户编号
	Balance   int64         `bson:"balance"`   // 余额，单位分
	Credit    int64         `bson:"credit"`    // 授信金额，单位分
	Status    AccountStatus `bson:"status"`    // （1=正常 2=禁用 9=注销）
	CreatedAt time.Time     `bson:"createdAt"` // 创建时间
	UpdatedAt time.Time     `bson:"updatedAt"` // 更新时间
}

// 储值账户状态（1=正常 2=禁用 9=注销）
type AccountStatus string

const (
	ACCOUNT_STATUS_NORMAL    AccountStatus = "1"
	ACCOUNT_STATUS_FORBIDDEN AccountStatus = "2"
	ACCOUNT_STATUS_INVALID   AccountStatus = "9"
)

// 现金账户交易：ms_cash_trade
type CashTrade struct {
	TradeId   string     `bson:"tradeId"`   // 交易编号
	Amount    int64      `bson:"amount"`    // 交易金额，单位分
	Remark    string     `bson:"remark"`    // 交易备注
	AccountNo string     `bson:"accountNo"` // 账户编号
	Balance   int64      `bson:"balance"`   // 交易后余额
	Type      *TradeType `bson:"type"`      // 交易类型：11=账户充值  12=账户提现  21=账户消费 31=转账转出  32=转账转入
	CreatedAt time.Time  `bson:"createdAt"` // 创建时间
	UpdatedAt time.Time  `bson:"updatedAt"` // 更新时间
}

var (
	TRADE_TYPE_RECHARGE *TradeType = &TradeType{Code: "11", Msg: "充值"}
	TRADE_TYPE_EXTRACT  *TradeType = &TradeType{Code: "11", Msg: "提现"}
	TRADE_TYPE_CONSUME  *TradeType = &TradeType{Code: "11", Msg: "消费"}
	TRADE_TYPE_ROLLOUT  *TradeType = &TradeType{Code: "11", Msg: "转出"}
	TRADE_TYPE_ROLLIN   *TradeType = &TradeType{Code: "11", Msg: "装入"}
)

// 现金账户交易类型，每个类型对应一个记录。 注意：交易类型与操作类型的区别，交易类型对应账户余额的加减，操作类型对用用户的操作
// 交易类型：11=账户充值  12=账户提现  21=账户消费 31=转账转出  32=转账转入
type TradeType struct {
	Code string `bson:"code"`
	Msg  string `bson:"msg"`
	// 账户充值扩展字段,记录充值方式、充值渠道等信息

	// 账户体现扩展字段，体现卡号、提现渠道

	// 账户消费扩展字段，消费订单号，消费金额，退款金额
	OrderId       string `bson:"orderId"`
	ExpenseAmount int64  `bson:"expenseAmount"`
	RefundAmount  int64  `bson:"refundAmount"`

	// 转账转出扩展字段，转出目标账户编号
	TargetAccountNo string `bson:"targetAccountNo"`
	// 转账输入扩展字段，转入来源账户编号
	OriginAccountNo string `bson:"originAccountNo"`
}

// TCC操作记录： ms_cash_tcc
// 微服务事务TCC解决方案
type CashTcc struct {
	TradeId     string     `bson:"tradeId"`   // 交易编号
	Status      string     `bson:"status"`    // 操作状态 （0=预操作 1=确认 2=取消 3=异常）
	Message     string     `bson:"message"`   // 处理结果说明
	Trade       *CashTrade `bson:"trade"`     // 交易记录，确认时更新交易记录，取消时不做处理
	CreatedAt   time.Time  `bson:"createdAt"` // 创建时间 ，超过1分钟还是预操作状态的记录，发送异常报警，人工干预
	ConfirmTime time.Time  `bson:"confirmAt"` // 确认时间
	CancelTime  time.Time  `bson:"cancelAt"`  // 取消时间
}
