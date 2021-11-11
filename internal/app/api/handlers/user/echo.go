package user

import (
	"context"
	"encoding/json"
	"github.com/osamingo/jsonrpc/v2"
	export2 "myjsonrpcv2/internal/app/api/export"
	types2 "myjsonrpcv2/internal/app/api/types"
)

const methodRouteNameEcho = "Echo"

type EchoHandler struct{}

var _ export2.Handler = &EchoHandler{}

func (h EchoHandler) Name() string {
	return methodRouteNameEcho;
}

func (h EchoHandler) Params() interface{} {
	return &types2.EchoParams{}
}

func (h EchoHandler) Result() interface{} {
	return &types2.EchoResult{}
}

func (h EchoHandler) ServeJSONRPC(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc.Error) {
	var p types2.EchoParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return types2.EchoResult{
		//Message: "Hello, " + p.Name,
		Message: `作家：陈少梅

Artist: Chen Shaomei

作品名称：《雪霁图》荣宝斋1:1限量版精品复制

Title：Rong Bao Zhai Limited Edition of Excellent 1:1 Duplicate Xue Ji Tu

跨链编号：IDA#202115#00171

Work No.: IDA#202115 #00171 

实物规格：94.5x46.7cm

Specifications: 94.5x46.7cm 

实物材质：设色纸本

Material：Ink and Color on Paper 

装裱：镜心

Mounting：Heart Mirror

全球限量发行10000幅，本幅编号为#00171。

With a limited edition of 10,000 pieces worldwide, this piece is numbered #00171.

作品简介：《雪霁图》是陈少梅先生于1948年所作，他以北宗为体，以南宗为用，以北宗蓄其势，以南宗添其韵，博采众长，笔墨潇洒而不失法度，艺术面貌渊穆清华，洋溢着和谐的节奏与优雅的韵律，将清逸淡雅、寓柔秀于阳刚的清劲画风付诸于纸上，犹使人置身于画中的景象。荣宝斋制作的本批作品，来源于荣宝斋馆藏精品1:1复制，保证原作的真实性。采用非国标印刷、荣宝斋独有工艺制作，全球限量制作、发行10000幅，所有作品由荣宝斋出具防伪证书。

Artwork Introduction: Xue Ji Tu was completed by Chen Shaomei in 1948. Taking the Northern School as the form and the Southern School as the specific technique, he draw the general structure with the northern sect, and added the specific charm to the southern sect. He drew on the strengths of others, with his techniques being chic but not out of order, and his artistic style being elegant, high and magnificent. He was permeated with harmony and graceful rhythm. He put the elegant, fresh and strong style of painting on paper, showing the femininity in the masculinity, which put people in the scene of the painting. This batch of works produced by Rongbaozhai are the 1:1 reproduction to ensure the authenticity of the original works. It is produced by non-national standard printing and the unique process of Rongbaozhai. 10,000 pieces are produced and distributed in limited quantities worldwide. All works are issued with anti-counterfeiting certificates by Rongbaozhai. 

作家简介：陈少梅（1909—1954），名云彰，又名云鹑，号升湖，字少梅，以字行。生于湖南衡山的一个书香之家，自幼随父学习书画诗文，深受中国传统文化的熏陶。15岁加入金北楼、陈师曾等发起组织的“中国画学研究会”，17岁成为名噪一时的“湖社画会”之骨干，22岁主持“湖社天津分会”，成为津门画坛领袖。1930年他的作品获“比利时建国百年国际博览会”美术银奖，以后开始在画坛崭露头角，成为京津一带颇有影响的画家。新中国成立后，他任中国美术家协会天津分会主席、天津美术学校校长。

北京匡时2017年秋拍，陈少梅《西园雅集图》以1300万元起拍，最终以2530万元成交。

Artist Introduction: Chen Shaomei(1909-1954), naming himself Yunzhang with the art name as Shenghu and the courtesy name Shaomei, was only called his courtesy name by others. Born at a literary family in Hengshan, Hunan province, he learned calligraphy and poetry under the instruction of his father since young, and was profoundly nurtured by Chinese traditional culture. At the age of 15, he joined the Chinese Painting Research Association launched and organized by Jinbeilou, Chen Shizeng, etc. He became the backbone of the once famous Hu She Club at 17 and headed the Tianjin Branch of the club at 22. In 1930, his work was granted the Silver Award of fine art in the International Expo for the Centennial of the Founding of Belgium. Later, he began to distinguish himself in the art world and became an influential painter in the Beijing-Tianjin area. After New China was founded, he took over as the chairman of the Tianjin branch of the Chinese Artists Association and the principal of Tianjin Fine Arts School.

In the 2017 autumn auction of Beijing Council, Chen Shaomei's Picture of Xiyuan Gathering started at 13 million yuan and was sold for 25.3 million yuan in the end.

参考资料：陈少梅在继承传统方面是集大成者，他心怀虔诚地对待所有优秀绘画传统，擅长山水、人物、走兽，工写兼长，传统功力深厚。尤其是他的人物画更为人所推重，张大千先生曾称赞陈少梅的人物画“衣纹有宋人风格”。

References: Chen Shaomei was a master in inheriting traditions. He treated all outstanding painting traditions with pious heart and was good at painting landscapes, figures, and animals. He performed excellently in meticulous painting or realistic painting, which showed his very solid traditional painting skills. His figure paintings were especially highly praised. Chang Da-Chien once praised Chen Shaomei's figure paintings for \"the texture of clothes has the Song Dynasty style\".
  
购买注意事项：  

Precautions Before Purchasing:  

1.用户购买后，可选择通过WWW.IP.PUB实物提货，实物将由中国发出。  

1.After purchase, users can choose to pick up the goods via WWW.IP.PUB, which will be sent by China.  

2.复制品与复刻品的定义、区别？  

2.What's the definition and distinction between copy and replica?  

注：复制品、复刻品、仿真复制、高仿书画都属于常见习惯用语。目前市场中常见的国标技术一般为“珂罗版”和“数字微喷”技术，均属于印刷技术。  

Note: Copy, replica, simulation replication and counterfeited calligraphy and painting are common idioms. At present, the common national standard technology in the market is generally “collotype” and “digital micro spray” technology, which belong to printing technology.  

3.本次荣宝斋复制品/复刻品属于哪个标准？  

3.What standard does this Rongbaozhai copy / replica belong to?  

荣宝斋出品的此批作品，均为馆藏精品，进行1:1复制。能保证原作品的真实性，在制作过程中可以与原作比照制作，不同于市场中只有电子版的印刷品，失色、失真的问题较为严重。  

This batch of works produced by Rongbaozhai are all high-quality works in the collection, which are copied 1:1. To ensure the authenticity of the original work, it can be compared with the original work in the production process. Different from the printing products of the electronic version on the market, the problem of color loss and distortion is more serious.  

4.荣宝斋如何保证本批作品不增发？  

4.How does Rongbaozhai ensure that this batch of works will not be issued additionally?  

荣宝斋将提供限量发行授权书，并注明了限量数量。  

Rongbaozhai will provide a limited distribution authorization letter, and indicate the limited number.  

5.本批次作品是否具备荣宝斋官方授权？  

5.Does this batch of works have the official authorization of Rongbaozhai?  

此批书画的制作，均由荣宝斋特批，资料室提档，转由制作部门制作，每张画都将具备荣宝斋出具作品证书。  

The batch of painting and calligraphy is approved by Rongbaozhai, filed by the archive room file and produced by the production department. Each painting will have a work certificate issued by Rongbaozhai.`,
	}, nil
}
