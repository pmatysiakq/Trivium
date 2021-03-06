package TriviumGo

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

// Initialize test vectors for basic encoding test
var(
	expectedKeyStream1 = "38EB86FF730D7A9CAF8DF13A4420540DBB7B651464C87501552041C249F29A64D2FBF515610921EBE06C8F92CECF7F8098FF20CCCC6A62B97BE8EF7454FC80F9EAF2625D411F61E41F6BAEEDDD5FE202600BD472F6C9CD1E9134A745D900EF6C023E4486538F09930CFD37157C0EB57C3EF6C954C42E707D52B743AD83CFF2979A203CF7B2F3F09C43D188AA13A5A2021EE998C42F777E9B67C3FA221A0AA1B041AA9E86BC2F5C52AFF11F7D9EE480CB1187B20EB46D582743A52D7CD080A24AEBF14772061C210843C18CEA2D2A275AE02FCB18E5D7942455FF77524E8A4CA51E369A847D1AEEFB9002FCD02342983CEAFA9D487CC2032B10192CD416310FA4"
	TestVector1 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"80000000000000000000",
	}

	expectedKeyStream2 = "61208D286BC1DC431171EDA5CAF79D9560B18ACEF26484417B651A47A3F7A80353F79AF8656DA4301A5E5A02E04265B182C67F5891220349F8CD1CD06597B77E242608D58B23D480E65A8957F3FA794F53802938517E00F63ACFB5EB6BD9EF468BBF3E25280DBA37FD0B0FDA76680A5596FF5271210EAFB170F3517238132C6F61ABC3A8B08C692F171CD714D9D15E8888F71F744EE561D289CD3180AE617FA43C81C882D7A946B0DDE1F00A6790E83F8641FB4573F75836F8E8397EF4A99F92178B25236ED340462A30E65B5AEC541A314B9D7D053106DA3E6E7D9AD15911A2DAE813763536B17B689113F870DC8EA1D1C455B4883DCB30227422946513157E"
	TestVector2 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00400000000000000000",
	}

	expectedKeyStream3 = "C8F9031DABF8DB03FF120D05512B5F24EAEA1BAB43201A5E93BF17F628E5B216D58577112F581A67DD5F962484ED4AC59202BA3509A73E119680B562F86DF0DCC26A443E26697FA16C0FF3D9152075BA1E81B900D609087F773FF30F4271F96E1BBABB533108766074D2C601E2C272194B64F775327F7C2250F7FCC0E377761FD0ECE1B05F3199C4AA3CE4BC21A912BC4488CC64EAA183C46AFAC7E3C07E459FCB08A75F1739D9AE1C489AAF307591D72E4E97339CC711FAA91AD16CB2D830B6DA6B9F7AD2563E7D14693343FCBBF5D59F85D5343A3DEA66CF2A81820881703A2F37D81A42F8A51B1977CECE86E477EB63E932B79F99B1241657B49953588760"
	TestVector3 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00002000000000000000",
	}

	expectedKeyStream4 = "F7E523040E86EA2C46A2BE705BFC62597F77E4649C0E71D51B288EDD4FC169BCD681F4603E192A7971E73290133E1E32F916D98B0D77F37927E1215C1D6AE0372128A7909071468084923FE0ECF981C967E59F15740C9A477EA468B81D0D82957ABD743AD921F9C106BEBC3A0092D4911BC68FAEE7BD0D72B9268640E040B6436BB914FF05BA243BC9030BA267C00AF52A0B7087AA05A5CAE114F89F0B0F8D926DAB38297B24360A8E9E77BA0C5F5D710AD1B28F2556874546C973A3A07FF09472948944739208A65AC2D1B05DA2CBE5999C536C06331D91081AF737BE28D6AAFCB4D7EDE955290D1E21038E1737E9E30A7681180BA89C746B77C2097D3CFD24"
	TestVector4 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000010000000000000",
	}

	expectedKeyStream5 = "4B430BDE0F574C7DE06E6A1918BFB4DEC0E2836071EB446C593EE1F2594533272E720E2A27992730E67D509EDF7BB0E62AEA85ED87B998FA6F53A0B77D26BBA2F2422BB51A6EDC8B05CB5C4A86C0FA2A8A631A6CA762075121400682E6D4198F33C021AF3F6EBA20C6CA1B705594806228788A3AC73641081FACD4EFF39FCEA5087CD26F0DB13D0ABE7946C81C6F1230DA00E6120E894143D63E01C76B4DE15B0128870032B23995F47A3C8421551E9F9DCB326394BEC319DAF4AF828F5A10861075D79B7886F8683D8126203AAB97DFCAA776F2F2B9D34F19EB09AE82F1193E38C6483D5C5A2F3B4F0FEF9559FE491E42995B6C9C0544D0EE71B45199257C2C"
	TestVector5 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000000080000000000",
	}

	expectedKeyStream6 = "4EAC0C5C7AD327084CFED7EAF72F6EB7FF20E11C65DFC1470C1D2EFAFCE2B1FA3ED6AFBE7CBEB677DB1189CB6892E81093B16FDC34199D0A26B89F06C86AC9D78E16BFC093187AA1F846233A522292CC2C1F829FE186613ACEBE8BBEF00470807A086475CD3A9F7A024DF5192B2A2C5E3599EA4BCFC4424C97F4848AAEE3CE620440260D186A7F351C3BA39ED04F7D3885EEB80873D8A7E0DB3EC9A85D157A6475FFEA67617555A9D1EC8C860C3F5632EB09543087C66E40061721F56E44330474CCF961E4D4508382AAFB7AC7246A1A4FAE1A2722404ED45A173C60E2785F06669914085078DDFF696B1EC62C314952519F59D3F83B6E304FDCA2BE1377DCB9"
	TestVector6 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000000000400000000",
	}

	expectedKeyStream7 = "2D97F227F2463F271F853BA10806888F8BA5733564557D5174A16172BD89D7E32ADF9C5B257DCC1693AB7CA6248F7A33D311FAD4D89D1C6EB7FCF8896D94E79C9C38E5671BBFF1255B0E0507AEA04A534CB9DF58FBFBEDF35A8D3E9ABF2D3575CFA09D25F4B65958F6335F2EFD5B6B26BD32854D2DCFE3BB9B8E45662E82B7552241EE8EF8C8E1D5F92660D4308EF4A2E049E8BC39C3E97009872036DADD6FC011236F68D2E53955353EF0BDF8C38AE589113090AC7B2668FF6D6430E3D86A8025D0F4B24C08E64008B664324A487BE95330E380025E1688B7E094120CB4EAA466C32BAC66FE0EB407FB61B2EF1D5EAC2D2CF69EB18D0CB218304BFB765A710B"
	TestVector7 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000000000002000000",
	}

	expectedKeyStream8 = "9EB8F6BD37474B5C2AC01BA1B3EEA5E1FBD4D4D1AE63EAFD81A4D2C900B5931273B37820BC68FCD69ED18F1ED4EC9334F39FEB330BAAEDEF2A1E51B218B385C8BF9DEDA5A5E009286220755F2C900A1F7FCD3B8C694FA76D82E19A4E598145AD064E7F863A6E7477E626D5736FEA8CEF48F0B8C6646C2CA76F14EC6864E9AD8A4137E7D79FC2569ABAC57B2740AB317297A07F2AD06CEC87C30BB474D86FC64E9E9149F5A9135DBF5F55D7811BD50611B11F92049D496D498353A1D60F676517A45AC8C930B8F3A12A1F742E06950A305F1DB8F4AE9EC0EC2821337843D25CCD3CFB3592FA440424F89A5CA600A1DD603A22C8913E34457D721B93A07D006BB0"
	TestVector8 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000000000000010000",
	}

	expectedKeyStream9 = "5D492E77F8FE62D769C6A142056BE9361FA0ADD8A54601DE615EBC04C4F8B2C12A8ED2DC9AB286A0F6C49C7AB319BA6AAFAAF0CD42D0A44C7DACBC90791855D8DF1884141AB121E7459DE30B2A0C85CDA0016453D350EF6220526CE33C1806E3831BEE5BF226E560F165FFC3585F2E54F4845419893265CD88BE8CB302B20345FEEF3646A5600D104A2DE2ABAF9B41A997FEC3EAF10C2F50EBF04D12CE139F6EBF7544F9BC85FA776E266C368C249D274DF3ECA2CE43F7A611AF0047D22E86012932E68122A3BD81A5EE6922FA11067F5DA90D30B775F752831AE09527B72DBEFC4DDA52DA28143720A31DFE662CDB73DED478E8541E3472AB4F80712250FE41"
	TestVector9 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000000000000000080",
	}

	expectedKeyStream10 = "FBE0BF265859051B517A2E4E239FC97F563203161907CF2DE7A8790FA1B2E9CDF75292030268B7382B4C1A759AA2599A285549986E74805903801A4CB5A5D4F20F1BE95091B8EA857B062AD52BADF47784AC6D9B2E3F85A9D79995043302F0FDF8B76E5BC8B7B4F0AA46CD20DDA04FDD197BC5E1635496828F2DBFB23F6BD5D080F9075437BAC73F696D0ABE3972F5FCE2192E5FCC13C0CB77D0ABA09126838D31A2D38A2087C46304C8A63B54109F679B0B1BC71E72A58D6DD3E0A3FF890D4A68450EB0910A98EF1853E0FC1BED8AB6BB08DF5F167D34008C2A85284D4B886DD56883EE92BF18E69121670B4C81A5689C9B0538373D22EB923A28A2DB44C0EB"
	TestVector10 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"00000000000000000000",
	}

	expectedKeyStream11 = "AB97616E7BAF0921F424B2573BFA15BDCA01898ABBE6AB77279AB1D732ABD10512769CC69FAB34E03B807F92C96627C17656BCC9BD9D377240A2B6FDBD7844538B6043B0C76E991D352361045F1B072456C04D7972CEE60F8727798497D4D3AC69CE6781D09795DDCD16A48236785201AF1884E69825977D4988B6C45409E187EDF6242C03842CAC95579088A4B6C60496FC3C80A17BB01FD25F647D1998DE1EB70FB1D5FE3754DD7C8801576C7CB6E06CBEE317B091C542BA91B83B1389CE50EC087DE06F92F49E8B0C69BCE33276CC3EC0B0472494E4DE3F8EBE37A9599EC8B6F46FCB72459FE7C273B3EB5188A6FC42B7733D74F7D2462B14D34EC0E24971"
	TestVector11 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"09090909090909090909",
	}

	expectedKeyStream12 = "AD35E17B1971AF6B5B3E365FA64EB4CF7EBF023D520889AD71F1A07AA2E0FF44CEE32D09CF77360C52434D462B53A405EF5A60D82A0F5E3CF3321B3727A3D61E850E4A9D5D6CE76AB58B97B6BD3F965419FA993BCE0C618B0200A0B024B6EF301BD3E32B173F8D63D2077887DE45D2EE842B71B407CCD7629EB5185732F048B0360B6895BCDBBB111156A9F9892850F0F059EF19E4EB265896AB3FD71017E156D80071542C5EB8843294053EEF33E6C0E62C640182BD26D0CB0677422964E23B3889B97439DBBAF76E480EEDA6EB387101DD97679826677500297211E7BC71218E4CC353596562C17E1B7BB4228B9869149CE108339442C16EE2A4E6A20EC1EA"
	TestVector12 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"12121212121212121212",
	}

	expectedKeyStream13 = "575AC77DCD54E7F48D837ADE0A88C70D777839B310C297835F5BB88C1A5FFD1569A4C4676E7D7CEA864FA1AD2F78E3E044E7D145827DD6E4095CB308111F5C720961EB35FD322EAE8011052513323CC51B0F4EBFBC25104A2E947B98B545D00DAE223FFF25AAAE27A6AD25CC7CBAD627559664E720CFB9061DD15F8C6BAE45F8ED7100988973CC712992DB2C42BA0FDF3F0A5B49B41398E96214F1C2F4F1C0EFEC17C6A8E80875D3EB40E973710614E572233CC2D88AABBA1B67D40365A10641EBA754FCAC87029BD9D475FB79FBD6FD1ADE18D75DF85CDF875C517862855665C6DD45F3999032410D7B71EFDE3E4FDAA923E9346121C8EA2EC999575870E043"
	TestVector13 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"1B1B1B1B1B1B1B1B1B1B",
	}

	expectedKeyStream14 = "5502DF8070387D1237EF5213B5F19EB79B9CC30810DB966DD2B25C9249C0378B17423F4E788BC3E82FBFC7FA3FED4A5704BA35CF9CBD12DCFD56CB6DD431A4D5A68899B41A40E33776B5F06905D417267698EB1306923997588459F83620CFEEC503FB9299F7640B67554079E04A404FF84966D5838D35D639065CE9B63BCE001D0DF339EEB936F041748F00AD03506551E81724ECCC589A0848C24962EF1766008B90259738515712B2F23EE6C81E3985439A1D6E9E2DBD6CE0207DA4ACFAE4D3D23D51D0786C1724E9043C2FAF83E632292A3640B2CE6C129B509EFC9641B62BED08E57E2E3FF6F19B335B49468CB5D49DD8CA4575617BEE7A3A2B12DFF7D4"
	TestVector14 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"24242424242424242424",
	}

	expectedKeyStream15 = "FC87849624217531385850ABE1CA16D5792A45F8FC40638B0BDFB7A32D5B53CC6751FE1CBE1BA673D2113AB900CFC9C3095CA85EAAF17DC617249FEB362A3F439DFDFB4E65756A17EB25E7780B76A5E38387FD2E6680F1C95E4316889062676386883BFCB70CF7D3B996ADE723A640F06F4FB89998DC89836BE5DE6867AB009FE994EA686CD3B38CBAE1623243671723E73659DB9E9AE918738B9178EEFB9165A63359C35F56DA951C50C8507CB759223AB159C8330DCDC1E96B59840B4A7F874A1497F3FF052FD845DA746EA324C700A10BD0C4E66F5E542542571EF612FABFC71F6CB099F615D0D5A09BDBD566337E40A25B7F2665C28CE0DC8ACCBB13F76E"
	TestVector15 = Trivium{
		Iv:		"00000000000000000000",
		Key: 	"2D2D2D2D2D2D2D2D2D2D",
	}

	expectedKeyStream16 = "F8901736640549E3BA7D42EA2D07B9F49233C18D773008BD755585B1A8CBAB86C1E9A9B91F1AD33483FD6EE3696D659C9374260456A36AAE11F033A519CBD5D787423582AF64475C3A9C092E32A53C5FE07D35B4C9CA288A89A43DEF3913EA9237CA43342F3F8E83AD3A5C38D463516F94E3724455656A36279E3E924D442F06D94389A90E6F3BF2BB4C8B057339AAD8AA2FEA238C29FCAC0D1FF1CB2535A07058BA995DD44CFC54CCEC54A5405B944C532D74E50EA370CDF1BA1CBAE93FC0B54844151714E56A3A2BBFBA426A1D60F9A4F265210A91EC29259AE2035234091C49FFB1893FA102D425C57C39EB4916F6D148DC83EBF7DE51EEB9ABFE045FB282"
	TestVector16 = Trivium{
		Iv:		"80000000000000000000",
		Key: 	"00000000000000000000",
	}

	expectedKeyStream17 = "ACBB386876653D15010DEFA7C65B36D701CFAF927B417550BE32D0444A24DEB589159B965C6740823F6BDFC378174AE2F664DCA0B68C621D2775BD13E6A788DFA322DB3314E80834F573DE2CF1AD6344D39AECB3FD0D35FE0CD5379ADA39A7531BB8C0B0B2C54DD1567C9AD9714D6719B1678401845620A86E79304946162C55B2799DD9D2DBEDEF0892A7A11BDFAE77B8F0E08BF83FB779653A190C295DB320C1BA5DC9BA40ADF8535F6C761F5E4D3393655B871116C22D14894C1A503361FED613722688EF02F23DE38710FBD2DCB80322A0C5ED64A34B7FFB1474B11D64076326F16FEAA3AB57401625E5C1614E9C162E69003033948566F527E732BEBD2B"
	TestVector17 = Trivium{
		Iv:		"00400000000000000000",
		Key: 	"00000000000000000000",
	}

	expectedKeyStream18 = "88BD48945DEA0BEB94D1F13FC589D61F4961046D4054B2EC274709DB1D8CC5472D1CD07D3CEBFC31E56DFED58029E598FB45D1954B6C86C9CC5EF422FFADFE324E127B792FE08EB7D40FC4AF45DF47E7D9C95265A45B30E9C49924A357CE34F621E7011BDFC11A1A4562F7329C90DE972BCCE296E347AD60F3167232A0664A8FA5ED498A0B8074D6E074135FABF441E00F7E8254A3D36DEBFFCFB257093BCC4E4CA597CF909A2CF207B7D34B8AAE33608BFAD0F9C3B66E9888356B981453A56826CE6C8180F9A29D278AE64B2C1219E3AB6337DEA7E488C50F6E5A7A29437716D0B918AAC8F3B07271977A6AAC0A76306C8B741EE7AB6B937AE7F0A834E4494F"
	TestVector18 = Trivium{
		Iv:		"00002000000000000000",
		Key: 	"00000000000000000000",
	}

	expectedKeyStream19 = "890782471E32E042C14767285A9BBD89605FEE5E38B9E78E3D750821AC7B4864A28DA27EB2EBBD6413CC6A5066E5240506E3F37C22876A7E9557C6B1BE1CE300F3679559680368D142E786B64F72D83C61F621532E8C4DE06A4934F643F38EC86B3524495D4682ABE7AFBC58BE5CA4354FC41C551D23DCD71FCE134A34B859FCB557DFB21905F9DB8430D1DF5089DF091B8B9C12C43B7A73E749EC87E88EBB42E562835D87C84E8B70652B0EEB6408E83F5673197C4786E1F2BBE65FB1CA3D9C190270788FC618E279E2B3EB94360F4FF1D8D15A2608D868AEE21E3DF72439D2B6B032EBAFA3F6513708506452B93C9C045678BF7C4853721FB41016B3362F9B"
	TestVector19 = Trivium{
		Iv:		"00000010000000000000",
		Key: 	"00000000000000000000",
	}

	expectedKeyStream20 = "27882DB265B89E59AD89611880FE6D125A4E98744C91CBFE1706E65A0A00A1EE3C020004AA287ECC209EA77650459840D73BB289482C732A024941E72E3244F0152FE0F80B054F3095C897D934FDA17A2230A4EC72C94C4C7D056DA5296E9F40DC42D74BF2CA2005EF38058A7904B8555BED47C01F81C94EFF7F77A9CE60C7169A785C9D3600F86DECAAAAEDEAAA2E0E34EF24ACA194C7E972C4AC2B8FED1F5289C6E4DC658259C497B240CCEBC9F8E5EC981A64BA731252431655F9B3E863EE63CFADD85A742D79880BFA8BBF5399ACFD683834092A2C9C64063F11BE89B541EC396ADEC8632F7AFA35B58F5FFCEB0529D4E5A388075CB97FA383001B7D4746"
	TestVector20 = Trivium{
		Iv:		"00000000000400000000",
		Key: 	"00000000000000000000",
	}
)

// Initialize test vectors for advanced encoding test
var (
	expectedKeyStream21 = "F4CD954A717F26A7D6930830C4E7CF0819F80E03F25F342C64ADC66ABA7F8A8E6EAA49F23632AE3CD41A7BD290A0132F81C6D4043B6E397D7388F3A03B5FE358C04C24A6938C8AF8A491D5E481271E0E601338F01067A86A795CA493AA4FF265619B8D448B706B7C88EE8395FC79E5B51AB40245BBF7773AE67DF86FCFB71F30011A0D7EC32FA102C66C164CFCB189AED9F6982E8C7370A6A37414781192CEB155C534C1C8C9E53FDEADF2D3D0577DAD3A8EB2F6E5265F1E831C86844670BC6948107374A9CE3AAF78221AE77789247CF6896A249ED75DCE0CF2D30EB9D889A0C61C9F480E5C07381DED9FAB2AD54333E82C89BA92E6E47FD828F1A66A8656E0"

	TestVector21 = Trivium{
		Iv:		"0D74DB42A91077DE45AC",
		Key: 	"0053A6F94C9FF24598EB",
	}

	expectedKeyStream22 = "A850A970ABCF5F73BCC5DB76F6B5E856362F1B36AC498D05C20FBE7763598DE1FD98B03CC54060E8C9C19B16490C665C3636A03BAB46656A695ED75F0E659F04D5F687A689F19C4F2258E212C3FAD8BA68625155CC92A43F282BB0F2F6851D4B6E748CE6E0774590F9A6F38DD1DA63BBB1C977697A5CE4BABE127A1201AE7520FE6E274131503C78178A5678020DC3E2E151B621C9DFC0DE065EADB396A149CDFB92E0009ECFC50237C81B0067B07E0E794AF78468E88B9E472E723ABC73C21A85E6891E2A605CF81224112E596B40A68C9D971AFC4376220B8160DCB36D55C04B0C21AF68A2EC38B4145F32BEDCE26012CF2151D6768E4CD025114C5E3149A5"

	TestVector22 = Trivium{
		Iv:		"167DE44BB21980E74EB5",
		Key: 	"0558ABFE51A4F74A9DF0",
	}

	expectedKeyStream23 = "DE9410C5134BCBADD0D2D95684E838183B91B0E8C1FA173C38F5B75103ABF0B8546EDDE22D6BFB3BF1C0754C6C42982FAEA5A3BD03DAFC0D1586E389B78C587236A54BF14629F0AE515610B8A3CC4A032490073054049502F63B1FCAA162ADA236D0A478721FDF31AB756FAA89B037D9921C2D3425F497B601906DA210ED28D5C282E166344A3F9950C174BD17091FA3BBA2A033F9111F4CF9B9462C45B6F422C51B24E7DF00D74C7229A8FDCF2AF4281A6BF656A228674643E3BED94A6AB8CD87545070DC5A80929893998A3CEE17036E9B7C13B14009C863E26BFCA5512254B1BCCF52498CEFEA7E4FD4435124F33A674B3A358318720F1F86CA84FF157D04"

	TestVector23 = Trivium{
		Iv:		"1F86ED54BB2289F057BE",
		Key: 	"0A5DB00356A9FC4FA2F5",
	}
)

// Generate messages of zero bytes
var plaintext = GenerateXZeroMsg(1024)
var longPlaintext = GenerateXZeroMsg(262144)	// 2^18 0.13MB??

// TestTriviumEncoding is basic encoding test for short message
func TestTriviumEncoding(t *testing.T) {
	testCases := []struct{
		expectedStreamKey 	string
		calculatedStreamKey	string
	}{
		{
			expectedStreamKey:   expectedKeyStream1,
			calculatedStreamKey: TestVector1.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream2,
			calculatedStreamKey: TestVector2.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream3,
			calculatedStreamKey: TestVector3.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream4,
			calculatedStreamKey: TestVector4.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream5,
			calculatedStreamKey: TestVector5.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream6,
			calculatedStreamKey: TestVector6.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream7,
			calculatedStreamKey: TestVector7.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream8,
			calculatedStreamKey: TestVector8.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream9,
			calculatedStreamKey: TestVector9.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream10,
			calculatedStreamKey: TestVector10.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream11,
			calculatedStreamKey: TestVector11.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream12,
			calculatedStreamKey: TestVector12.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream13,
			calculatedStreamKey: TestVector13.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream14,
			calculatedStreamKey: TestVector14.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream15,
			calculatedStreamKey: TestVector15.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream16,
			calculatedStreamKey: TestVector16.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream17,
			calculatedStreamKey: TestVector17.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream18,
			calculatedStreamKey: TestVector18.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream19,
			calculatedStreamKey: TestVector19.Encrypt(plaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream20,
			calculatedStreamKey: TestVector20.Encrypt(plaintext),
		},
	}

	index := 1
	for _, testCase := range testCases {
		t.Run("Testing:" + strconv.Itoa(index), func(t *testing.T) {
			index++
			// Take only slices of bits to comparison [1:64], [192:319], [488:511]
			calculated 	:= testCase.calculatedStreamKey[0:128] + testCase.calculatedStreamKey[384:640] + testCase.calculatedStreamKey[896:]
			expected 	:= testCase.expectedStreamKey

			if expected != calculated {
				t.Error("Generated Calculated StreamKey (cipher for zero msg) is not equal to expected StreamKey")
				fmt.Println("EXP:", expected)
				fmt.Println("CAL:", calculated)
			}
		})
	}
}

// TestTriviumLongMsg is advanced encoding test.
// Here we provide very long message to encode.
func TestTriviumLongMsg(t *testing.T) {
	testCases := []struct{
		expectedStreamKey 	string
		calculatedStreamKey	string
	}{
		{
			expectedStreamKey:   expectedKeyStream21,
			calculatedStreamKey: TestVector21.Encrypt(longPlaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream22,
			calculatedStreamKey: TestVector22.Encrypt(longPlaintext),
		},
		{
			expectedStreamKey: 	expectedKeyStream23,
			calculatedStreamKey: TestVector23.Encrypt(longPlaintext),
		},
	}

	index := 21
	for _, testCase := range testCases {
		t.Run("Testing:" + strconv.Itoa(index), func(t *testing.T) {
			index++
			// Take only slices of bits to comparison [1:64], [65472:65599], [131008:131071]
			calculated 	:= testCase.calculatedStreamKey[0:128] + testCase.calculatedStreamKey[130944:131200] + testCase.calculatedStreamKey[262016:]
			expected 	:= testCase.expectedStreamKey

			if expected != calculated {
				t.Error("Generated Calculated StreamKey (cipher for zero msg) is not equal to expected StreamKey")
				fmt.Println("EXP:", expected)
				fmt.Println("CAL:", calculated)
			}
		})
	}
}

// TestTriviumLongMsg is advanced encoding test.
// Here we provide very long message to encode.
func TestEfficiency(t *testing.T) {
	testCases := []struct{
		msg 		string
		msgLength	int
	}{
		{
			// 2^3 HEX characters
			msg: RandStringRunes(8),
			msgLength: 8,
		},
		{
			// 2^4 HEX characters
			msg: RandStringRunes(16),
			msgLength: 16,
		},
		{
			// 2^5 HEX characters
			msg: RandStringRunes(32),
			msgLength: 32,
		},
		{
			// 2^6 HEX characters
			msg: RandStringRunes(64),
			msgLength: 64,
		},
		{
			// 2^7 HEX characters
			msg: RandStringRunes(128),
			msgLength: 128,
		},
		{
			// 2^8 HEX characters
			msg: RandStringRunes(256),
			msgLength: 256,
		},
		{
			// 2^9 HEX characters
			msg: RandStringRunes(512),
			msgLength: 512,
		},
		{
			// 2^10 HEX characters
			msg: RandStringRunes(1024),
			msgLength: 1024,
		},
		{
			// 2^11 HEX characters
			msg: RandStringRunes(2048),
			msgLength: 2048,
		},
		{
			// 2^12 HEX characters
			msg: RandStringRunes(4096),
			msgLength: 4096,
		},
		{
			// 2^13 HEX characters
			msg: RandStringRunes(8192),
			msgLength: 8192,

		},
		{
			// 2^14 HEX characters
			msg: RandStringRunes(16384),
			msgLength: 16384,
		},
		{
			// 2^15 HEX characters
			msg: RandStringRunes(32768),
			msgLength: 32768,
		},
		{// 2^16 HEX characters
			msg: RandStringRunes(65536),
			msgLength: 65536,
		},
		{
			// 2^17 HEX characters
			msg: RandStringRunes(131072),
			msgLength: 131072,
		},
		{
			// 2^18 HEX characters
			msg: RandStringRunes(262144),
			msgLength: 262144,
		},
		{
			// 2^19 HEX characters
			msg: RandStringRunes(524288),
			msgLength: 524288,
		},
		//{
		//	// 2^20 HEX characters
		//	msg: RandStringRunes(1048576),
		//	msgLength: 1048576,
		//},
		//{
		//	// 2^21 HEX characters
		//	msg: RandStringRunes(2097152),
		//	msgLength: 2097152,
		//},
	}

	file, err := os.OpenFile("../trivium-efficiency-2.csv", os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write([]byte("msg-len,time-elapsed,peak-memory\n")); err != nil {
		fmt.Printf("Couldn't write to file... Error:%v\n", err)
	}
	file.Close()

	file, err = os.OpenFile("../trivium-efficiency-2.csv", os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, testCase := range testCases {
		t.Run("Msg Length:" + strconv.Itoa(testCase.msgLength), func(t *testing.T) {
			startTime := time.Now()

			startMemory := PrintMemUsage("Memory allocation start: ")
			trivium := NewTrivium("54656c636f52756c6573", "4e696d6f6d706f6a6563")
			trivium.Encrypt(testCase.msg)
			stopTime := time.Now()
			stopMemory := PrintMemUsage("Memory allocation end: ")

			elapsedTime := stopTime.Sub(startTime)
			memoryConsumed := stopMemory - startMemory
			fmt.Printf("\tTotal time: %v\n", elapsedTime)
			fmt.Printf("\tTotal memory usage: %v bytes\n", memoryConsumed)

			line := fmt.Sprintf("%s,%s,%s\n", strconv.Itoa(testCase.msgLength), strconv.FormatInt(elapsedTime.Milliseconds(), 10), strconv.Itoa(int(memoryConsumed)))
			if _, err := file.Write([]byte(line)); err != nil {
				fmt.Printf("Couldn't write to file... Error:%v\n", err)
			}
		})
	}
}