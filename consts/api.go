package consts

// 淘宝API接口路径常量定义
//拼接授权 URL
//按照以下规则拼装授权URL，并引导业务同学访问：
//https://api.taobao.global/oauth/authorize?response_type=code&redirect_uri=${redirect_uri}&force_auth=true&client_id=${appkey}
//注意：请将「appkey」及「redirect_url」替换为您的开发者应用程式配置。
const (
	BaseApi = "https://api.taobao.global/rest"

	// 鉴权相关接口
	TaoBaoApiSignMethod          = "sha256"              // 签名方式：采用SHA-256算法进行签名
	TaoBaoApiRefreshAccessToken  = "/auth/token/refresh" // 刷新访问令牌接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193531.0.0.7b545288cD0hQP#/api?cid=8&path=/auth/token/refresh&methodType=GET/POST
	TaoBaoApiGenerateAccessToken = "/auth/token/create"  // 生成访问令牌接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193531.0.0.7b545288cD0hQP#/api?cid=8&path=/auth/token/create&methodType=GET/POST

	// 商品相关接口
	TaoBaoApiQueryAllProduct = "/product/get"            // 获取商品详情接口 https://open.taobao.global/doc/api.htm?spm=panama_open.panama_open_home.0.0.753123f9IuBxQs#/api?cid=14&path=/product/get&methodType=GET/POST
	TaoBaoApiProductSearch   = "/traffic/item/search"    // 商品搜索接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193494.0.0.3efe76976IxWtP#/api?cid=21&path=/traffic/item/search&methodType=GET/POST
	TaoBaoApiImageUpload     = "/upload/image"           // 图片上传接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193531.0.0.2d4a5288j6INTs#/api?cid=21&path=/upload/image&methodType=POST
	TaoBaoApiImgSearch       = "/traffic/item/imgsearch" // 以图找货接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193531.0.0.2d4a5288j6INTs#/api?cid=21&path=/traffic/item/imgsearch&methodType=GET/POST
	// 订单相关接口
	TaoBaoApiPurchaseOrderRender   = "/purchase/order/render"      // 订单预览接口 https://open.taobao.global/doc/api.htm?spm=panama_open.panama_open_home.0.0.753123f9VwZCVc#/api?cid=16&path=/purchase/order/render&methodType=GET/POST
	TaoBaoApiCreatePurchaseOrder   = "/purchase/order/create"      // 创建订单接口 https://open.taobao.global/doc/api.htm?spm=panama_open.panama_open_home.0.0.753123f9VwZCVc#/api?cid=16&path=/purchase/order/create&methodType=POST
	TaoBaoApiPurchaseOrderBatchPay = "/purchase/order/batch/pay"   // 批量支付接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193531.0.0.2d4a5288j6INTs#/api?cid=16&path=/purchase/order/batch/pay&methodType=GET/POST
	TaoBaoApiOrderCancel           = "/purchase/order/asyn/cancel" //取消采购单接口 https://open.taobao.global/doc/api.htm?spm=panama_open.panama_open_home.0.0.753123f9VwZCVc#/api?cid=16&path=/purchase/order/asyn/cancel&methodType=GET/POST
	TaoBaoApiLogisticsDetail       = "/logistics/detail/get"       // 查询采购单物流详情接口 https://distributor.taobao.global/apps/refund/list?spm=a1zlqa.report_list.apps_refund_list.d_btn
	TaoBaoApiOrdersQuery           = "/purchase/orders/query"      // 查询采购单接口 https://open.taobao.global/doc/api.htm?spm=a2o9m.11193531.0.0.2d4a5288j6INTs#/api?cid=16&path=/purchase/orders/query&methodType=GET/POST
)
