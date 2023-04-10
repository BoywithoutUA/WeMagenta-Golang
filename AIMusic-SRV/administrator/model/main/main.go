package main

import (
	"administrator/config"
	"administrator/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//此函数用于同步数据库表
func main() {
	Config := &config.AIMusicConfig{}
	v := viper.New()
	v.SetConfigFile("config.yml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(Config); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/AIMusic?charset=utf8mb4&parseTime=True&loc=Local",
		Config.MySQL.User, Config.MySQL.Password, Config.MySQL.Host, Config.MySQL.Port)
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: time.Second,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	},
	//)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true,
		//},
		//Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.Composition{})
	//_ = db.AutoMigrate(&model.User{})
	//_ = db.AutoMigrate(&model.Template{})
	_ = db.AutoMigrate(&model.Feedback{})
	//_ = db.AutoMigrate(&model.Usage{})
	//db.Create(&model.Template{
	//	Name: "《伤寒杂病论》",
	//	Type: 0,
	//	Detail: "东汉末年张仲景;" +
	//		"公元200年～210年左右;" +
	//		"伤寒杂病论》是中国传统医学著作之一，是一部论述外感病与内科杂病为主要内容的医学典籍，作者是东汉末年张仲景，是中国中医院校开设的主要基础课程之一，它系统地分析了伤寒的原因、症状、发展阶段和处理方法，创造性地确立了对伤寒病的“六经分类”的辨证施治原则，奠定了理、法、方、药的理论基础。;",
	//	Pic: "/AIMusic/pic/shzbl.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《难经》",
	//	Type: 0,
	//	Detail: "春秋战国时期扁鹊;" +
	//		"一般认为不晚于东汉;" +
	//		"《难经》之“难”字，有“问难”或“疑难”之义。全书共八十一难，采用问答方式，探讨和论述了中医的一些理论问题，内容包括脉诊、经络、脏腑、阴阳、病因、病机、营卫、腧穴、针刺、病证等方面。《难经》被认为是可以羽翼《灵》《素》的中医经典著作，其寸口诊法，对奇经八脉、三焦和命门的论述均为后世所继承。;",
	//	Pic: "/AIMusic/pic/nj.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《针灸甲乙经》",
	//	Type: 0,
	//	Detail: "西晋皇甫谧;" +
	//		"公元282年;" +
	//		"其书前六卷论述基础理论，后六卷记录各种疾病的临床治疗，包括病因、病机、症状、诊断、取穴、治法和预后等。采用分部和按经分类法，厘定了腧穴，详述了各部穴位的适应证和禁忌、针刺深度与灸的壮数，是我国现存最早的一部理论联系实际的针灸学专著。;",
	//	Pic: "/AIMusic/pic/zjjyj.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《脾胃论》",
	//	Type: 0,
	//	Detail: "元代李东垣;" +
	//		"公元1249年;" +
	//		"该书引用大量《内经》原文以阐述其脾胃论的主要观点和治疗方药。卷中阐述脾胃病的具体论治。卷下详述脾胃病与天地阴阳、升降浮沉的密切关系，并提出多种治疗方法，列方60余首，并附方义及服用法。所创补中益气汤、调中益气汤、升阳益胃汤、升阳散火汤等为临床所习用。;",
	//	Pic: "/AIMusic/pic/pwl.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《食疗本草》",
	//	Type: 0,
	//	Detail: "唐代家孟诜;" +
	//		"713～741年;" +
	//		"本书共载文227条，涉及260种食疗品。诸品名下，注明药性(温、平、寒、冷)，不载其味。正文述功效、禁忌及单方，间或论及形态、修治、产地等。首载菠薐、胡荽、莙荙、鳜鱼等食蔬。尤以动物脏器疗法与藻菌类食疗作用之记载引人注目，是一部内容丰富的古代营养学和食物疗法专著，对多数食物疗效和食用药品合理应用的阐述切合实际，至今仍有较高价值。;",
	//	Pic: "/AIMusic/pic/slbc.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《千金要方》",
	//	Type: 0,
	//	Detail: "唐朝孙思邈;" +
	//		"652年;" +
	//		"该书集唐代以前诊治经验之大成，是综合性临床医著，被誉为中国最早的临床百科全书。中首篇所列的《大医精诚》、《大医习业》，是中医学伦理学的基础；其妇、儿科专卷的论述，奠定了宋代妇、儿科独立的基础；其治内科病提倡以\"五脏六腑为纲，寒热虚实为目\"，并开创了脏腑分类方剂的先河。针灸孔穴主治的论述，为针灸治疗提供了准绳，阿是穴的选用、“同身寸”的提倡，对针灸取穴的准确性颇有帮助。因此，《千金要方》素为后世医学家所重视。;",
	//	Pic: "/AIMusic/pic/qjyf.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《铜人腧穴针灸图经》",
	//	Type: 0,
	//	Detail: "北宋王惟一;" +
	//		"1026年;" +
	//		"《铜人腧穴针灸图经》和针灸铜人是为是为首次国家级的经穴大整理，以十四经为纲，三百五十四穴为目，并附有插图十五副。全书编写体例统一，结构严谨。一方面集成了古代针灸著作的理论系统，另一方面便于临证取穴治疗和研究，为针灸的发展做出了很大贡献，堪称针灸学发展史上的里程碑。;",
	//	Pic: "/AIMusic/pic/trsxzjtj.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《肘后备急方》",
	//	Type: 0,
	//	Detail: "东晋葛洪;" +
	//		"326-341年;" +
	//		"《肘后备急方》中收载了多种疾病，其中有很多是珍贵的医学资料。这部书上描写的天花症状，以及其中对于天花的危险性、传染性的描述，都是世界上最早的记载。书中涉及到了肠结核、骨关节结核等多种疾病，可以说其论述的完备性并不亚于现代医学。书中还记载了被疯狗咬过后用疯狗的脑子涂在伤口上治疗的方法，该方法与近代巴斯德的狂犬病疫苗大体相近。对于流行病、传染病，书中更是提出了“疠气”的概念，认为这绝不是所谓的鬼神作祟，这种科学的认识方法在当今来讲，也是十分有见地的。书中对于恙虫病、疥虫病之类的寄生虫病的描述，也是世界医学史上出现时间最早，叙述最准确的。;",
	//	Pic: "/AIMusic/pic/zhbjf.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《小儿药证直诀》",
	//	Type: 0,
	//	Detail: "宋代钱乙撰;" +
	//		"1119年;" +
	//		"该书是现存很早的汉医儿科学著作，在儿科发展史上占有重要地位，是汉医儿科的奠基之作。由钱氏门人阎孝忠编集而成。本书是中国早期内容比较完整，并载有病案的儿科重要专著。书中简要地记述了小儿病的诊断与治疗，具有较高的临床实用价值。;",
	//	Pic: "/AIMusic/pic/jeyzzj.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《本草纲目》",
	//	Type: 0,
	//	Detail: "明代李时珍;" +
	//		"1552年-1578年;" +
	//		"本书虽为中药学专书，但涉及范围广泛，对植物学、动物学、矿物学、物理学、化学、农学等内容亦有很多记载，刊行后，促进了本草学的进一步发展，倪朱谟的《本草汇言》、赵学敏的《本草纲目拾遗》、黄宫绣的《本草求真》等，均是在其学说启示下而著成的本草典籍。达尔文其著作中亦多次引用本书的资料，并称之为“古代中国百科全书”。英国李约瑟( Joseph Needham)称赞李时珍为“药物学界中之王子”。本书为本草学集大成之作。;",
	//	Pic: "/AIMusic/pic/bcgm.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《温热论》",
	//	Type: 0,
	//	Detail: "清代叶天士;" +
	//		"不晚于清乾隆十一年(1746年);" +
	//		"《温热论》是温病学说的奠基之作。叶天士在《温热论》中提出的卫气营血辨证方法是十分彻底的创新，不仅放弃了六经的概念，而且放弃了表里的概念，完全从温热病的传变特点出发进行辨证。叶天士是一位十分杰出的临床医学家，对于温病经验极其丰富，卫气营血辨证正是在精熟临床经验基础上的高度概括，既简洁明了，又逻辑严密，抓住了温热病发展的几个关键性环节，具有极高的临床适用性，对后世影响深远。;",
	//	Pic: "/AIMusic/pic/wrl.png",
	//})
	//db.Create(&model.Template{
	//	Name: "《温疫论》",
	//	Type: 0,
	//	Detail: "明末吴又可;" +
	//		"1642年;" +
	//		"吴又可的 《温疫论》揭开了中医治疗外感热病的新篇章，他的温疫学说对其后的戴天章、杨栗山等医家均有重要影响，戴天章等人对吴又可学术思想进一步阐释，从而奠定了温疫学派的学术根基。《温疫论》是中医温病学发展史上具有划时代意义的标志性著作，是中医理论原创思维与临证实用新法的杰出体现。吴又可在《温疫论》中创立了表里九传辨证论治思维模式，创制了达原饮等治疗温疫的有效方剂。对后世温病学的形成与发展产生了深远影响。;",
	//	Pic: "/AIMusic/pic/wyl.png",
	//})
	//db.Create(&model.Template{
	//	Name: "扁鹊",
	//	Type: 1,
	//	Detail: "春秋战国时期;" +
	//		"公元前407-310;" +
	//		"《扁鹊内经》、《扁鹊外经》;" +
	//		"扁鹊奠定了祖国传统医学诊断法的基础。他用一生的时间，认真总结前人和民间经验，结合自己的医疗实践，在诊断、病理、治法上对祖国医学作出了卓越的贡献。扁鹊的医学经验，在我国医学史上占有承前启后的重要地位，对我国医学发展有较大影响。扁鹊是我国历史上最早应用脉诊来判断疾病的医生，并且提出了相应的脉诊理论。;",
	//	Pic: "/AIMusic/pic/bq.png",
	//})
	//db.Create(&model.Template{
	//	Name: "华佗",
	//	Type: 1,
	//	Detail: "东汉末年;" +
	//		"145年-208年;" +
	//		"《素问》、 《灵枢》、《难经》、《神农本草经》、《伤寒论》、《金匮要略》、《中藏经》;" +
	//		"华佗是中国历史上第一位创造手术外科的专家，也是世界上第一位发明麻醉剂“麻沸散”及发明用针灸医病的先驱者、创始人。“麻沸散”为外科医学的开拓和发展开创了新的研究领域。他的发明比美国的牙科医生摩尔顿( 1846年)发明乙醚麻醉获得成功要早1600多年。他所使用的“麻沸散”是世界史最早的麻醉剂。华佗采用酒服“麻沸散”施行腹部手术，开创了全身麻醉手术的先例。这种全身麻醉手术，在中国医学史上是空前的，在世界医学史上也是罕见的创举。;",
	//	Pic: "/AIMusic/pic/ht.png",
	//})
	//db.Create(&model.Template{
	//	Name: "张仲景",
	//	Type: 1,
	//	Detail: "东汉末年;" +
	//		"约公元150～154年—约公元215～219年;" +
	//		"《金匮要略》、《伤寒杂病论》、《伤寒论》;" +
	//		"张仲景公元205年写的医学著作《伤寒杂病论》对于推动后世医学的发展起了巨大的作用。他是我国历史上最最杰出的医学家之一。他为我国的医学发展做出了重要的贡献，也促进了我国医学的发展。他创作的传世巨著《伤寒杂病论》，确立的辨证论治原则，是中医临床的基本原则，是中医的灵魂所在，也是后学者研习中医必备的经典著作。;",
	//	Pic: "/AIMusic/pic/zzj.png",
	//})
	//db.Create(&model.Template{
	//	Name: "董奉",
	//	Type: 1,
	//	Detail: "东汉时期;" +
	//		"220年—280年;" +
	//		"无;" +
	//		"董奉医术的高明和不求名利、乐善好施的高尚医德被人们传为佳话，千秋流传。人们把他同当时谯郡的华佗、南阳的张仲景并称为“建安三神医”。后世以“杏林春暖”，“誉满杏林”称誉医术高尚的医家，唤中医为“杏林”。;",
	//	Pic: "/AIMusic/pic/df.png",
	//})
	//db.Create(&model.Template{
	//	Name: "皇甫谧",
	//	Type: 1,
	//	Detail: "三国西晋时期;" +
	//		"215年—282年;" +
	//		"《针灸甲乙经》;" +
	//		"其著作《针灸甲乙经》是中国第一部针灸学的专著，在医学史负有盛名，奠定了针灸学科理论基础，对针灸学以至整个医学事业的发展作出了不可磨灭的贡献。。他在针灸学史上，占有很高的学术地位，并被誉为“针灸鼻祖”。 现今的针灸医学不但在国内得到飞速发展，并且已经风靡世界，世界卫生组织已经正式批准，把针灸列为治疗专项，到处受到人们的欢迎。正因为如此，皇甫谧是我国古代历史上唯一与孔子齐名于世界文化史的历史名人。;",
	//	Pic: "/AIMusic/pic/hfm.png",
	//})
	//db.Create(&model.Template{
	//	Name: "孙思邈",
	//	Type: 1,
	//	Detail: "唐代;" +
	//		"541年—682年，存在争议;" +
	//		"《千金要方》，《千金翼方》，《明堂针灸图》，《千金髓方》等;" +
	//		"孙思邈对古典医学有深刻的研究，对民间验方十分重视，一生致力于医学临床研究，对内、外、妇、儿、五官、针灸各科都很精通，有二十四项成果开创了中国医药学史上的先河，特别是论述医德思想、倡导妇科、儿科、针灸穴位等都是前人未有。 他创立了从方、证、治三方面研究《伤寒杂病论》的方法，开后世以方类证的先河。孙思邈终身不仕，隐于山林。亲自采制药物，为人治病。他搜集民间验方、秘方，总结临床经验及前代医学理论，为医学和药物学作出重要贡献。后世尊其为“药王”。;",
	//	Pic: "/AIMusic/pic/ssm.png",
	//})
	//db.Create(&model.Template{
	//	Name: "王惟一",
	//	Type: 1,
	//	Detail: "宋代;" +
	//		"987-1067;" +
	//		"《新铸铜人腧穴针灸图经》;" +
	//		"王惟一“素校禁方，尤工厉石”,“创铸铜人为式”，考订经穴理论，为经穴理论的发展与规范化，以及针灸教学作出了巨大的贡献，是宋代杰出的针灸学家和医学教育家，为中国医学的发展做出了不可磨灭的贡献。王惟一所设计的铜人，对中国医学的发展，尤其在针灸学和针灸教学方面，起了很大的促进作用，故为历来针灸学家所推崇，即至现在仍有学习和研究的价值。;",
	//	Pic: "/AIMusic/pic/wwy.png",
	//})
	//db.Create(&model.Template{
	//	Name: "钱乙",
	//	Type: 1,
	//	Detail: "宋代;" +
	//		"约1032年－1113年;" +
	//		"《伤寒论发微》，《婴孺论》，《钱氏小儿方》，《小儿药证直诀》。现仅存《小儿药证直诀》;" +
	//		"钱乙由于对小儿科作了四十年的深入钻研，终于摸清了小儿病诊治的规律，积累了丰富的临证经验。现存《小儿药证直诀》，该书最早记载辨认麻疹法和记百日咳的证治；也是最早从皮疹的特征来鉴别天花、麻疹和水痘；记述多种初生疾病和小儿发育营养障碍疾患，以及多种著名有效的方剂；还创立了我国最早的儿科病历。此书一为历代中医所重视，列为研究儿科必读之书。它不仅是我国现存最早的第一部系统完整的儿科专著，而且也是世界上最早的儿科专著。;",
	//	Pic: "/AIMusic/pic/qy.png",
	//})
	//db.Create(&model.Template{
	//	Name: "李时珍",
	//	Type: 1,
	//	Detail: "明代;" +
	//		"1518年7月3日－1593年;" +
	//		"《本草纲目》《奇经八脉考》《濒湖脉学》等多种;" +
	//		"李时珍对发展中国药物学方面所做的卓越贡献是有口皆碑的。他在医疗方面，发展切脉诊断方面，都有着出色的成就。他将自己收集到的11096个单方、秘方、验方，各一一附录于各该药物之下予以论述，这也是他的一大创举。他研究中药学数十年，参考各种图书800多种，撰成《本草纲目》52卷，集明代药物学之大成。在诊断方面，他还撰有《濒湖脉学》一书，发展了中医诊断学。所著《奇经八脉考》一书，则是规范中医经络学说的一次有价值的努力。李时珍被誉为中国最著名的医药学家、世界著名的学者，当受之无愧。;",
	//	Pic: "/AIMusic/pic/lsz.png",
	//})
	//db.Create(&model.Template{
	//	Name: "杨继洲",
	//	Type: 1,
	//	Detail: "明代;" +
	//		"约1522年～1620年;" +
	//		"主要作品有《针灸大成》;" +
	//		"杨继洲是明代一位针灸学之集大成者，他总结了明末以前针灸学的重要成果，是继《针灸甲乙经》以后，对针灸学的又一次重要总结。《针灸大成》的问世，标志着中国古代针灸学已经发展到了相当成熟的地步，后人在论述针灸学时，大多将《针灸大成》作为最重要的参考书，这与该书的学术成就、所处的历史地位以及其对针灸学发展所作出的巨大贡献是分不开的。;",
	//	Pic: "/AIMusic/pic/yjz.png",
	//})
	//db.Create(&model.Template{
	//	Name: "李东垣",
	//	Type: 1,
	//	Detail: "金代;" +
	//		"1180年-1251年;" +
	//		"《内外伤辨惑论》、《脾胃论》、《兰室秘藏》、《医学发明》、《东垣试效方》《活法机要》等;" +
	//		"李杲通过长期的临床实践积累了一定的经验，提出“内伤脾胃，百病由生”的观点，形成了独具一格的脾胃内伤学说。是我国医学史上著名的金元四大家之一。是中医“脾胃学说”的创始人。同时，他还将内科疾病系统地分为外感和内伤两大类，这对临床上的诊断和治疗有很强的指导意义。.河间学派和易水学派为中国医学史上承前启后影响最大的两大学派，李杲为易水学派的中流砥柱。;",
	//	Pic: "/AIMusic/pic/ldh.png",
	//})
	//db.Create(&model.Template{
	//	Name: "万密斋",
	//	Type: 1,
	//	Detail: "明代;" +
	//		"1499-1582;" +
	//		"《万密斋医学全书》，《养生四要》;" +
	//		"万密斋的著作与学术思想源清流洁，本盛末荣，涉及儿、妇、内科及优生、优育、延龄、广嗣、养生、保健，堪称博大精深。回溯过去，因其承前启后，发皇古义，务实求是，颇多创见，对明清临证医学发挥了深刻的影响；瞻望未来，因其以人为本，方药齐备，施治灵活，实用性强，对当代中医药理论研究与临床工作也具有重要的参考价值",
	//	Pic: "/AIMusic/pic/wmz.png",
	//})
}
