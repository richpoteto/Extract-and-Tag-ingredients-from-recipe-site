package recipeingredients

var herbMap = map[string]struct{}{
	"ajwain":               {},
	"alligator pepper":     {},
	"allspice":             {},
	"amchoor":              {},
	"angelica":             {},
	"anise":                {},
	"aonori":               {},
	"aromatic ginger":      {},
	"asafoetida":           {},
	"basil":                {},
	"bay leaf":             {},
	"black cardamom":       {},
	"black mustard":        {},
	"black peppercorn":     {},
	"boldo":                {},
	"bolivian coriander":   {},
	"borage":               {},
	"brazilian pepper":     {},
	"brown mustard":        {},
	"bunium persicum":      {},
	"camphor":              {},
	"caraway":              {},
	"cardamom":             {},
	"cassia":               {},
	"cayenne pepper":       {},
	"celery powder":        {},
	"celery seed":          {},
	"charoli":              {},
	"chenpi":               {},
	"chervil":              {},
	"chili pepper":         {},
	"chives":               {},
	"cicely":               {},
	"cilantro":             {},
	"cinnamon":             {},
	"clove":                {},
	"coriander leaf":       {},
	"coriander seed":       {},
	"cress":                {},
	"cubeb":                {},
	"culantro":             {},
	"cumin":                {},
	"curry leaf":           {},
	"dill":                 {},
	"dill seed":            {},
	"dried lime":           {},
	"east asian pepper":    {},
	"epazote":              {},
	"fennel":               {},
	"fenugreek":            {},
	"fingerroot":           {},
	"garlic":               {},
	"garlic chives":        {},
	"ginger":               {},
	"golpar":               {},
	"grains of paradise":   {},
	"grains of selim":      {},
	"greater galangal":     {},
	"green peppercorn":     {},
	"hemp":                 {},
	"hoja santa":           {},
	"holy basil":           {},
	"horseradish":          {},
	"houttuynia cordata":   {},
	"hyssop":               {},
	"indian bay leaf":      {},
	"jimbu":                {},
	"juniper berry":        {},
	"kinh gioi":            {},
	"kokum":                {},
	"korarima":             {},
	"lavender":             {},
	"lemon balm":           {},
	"lemon grass":          {},
	"lemon myrtle":         {},
	"lemon verbena":        {},
	"lesser galangal":      {},
	"limnophila aromatica": {},
	"liquorice":            {},
	"litsea cubeba":        {},
	"long pepper":          {},
	"lovage":               {},
	"mace":                 {},
	"mahlab":               {},
	"mango-ginger":         {},
	"marjoram":             {},
	"mastic":               {},
	"mint":                 {},
	"mitsuba":              {},
	"mugwort":              {},
	"nigella":              {},
	"nigella sativa":       {},
	"njangsa":              {},
	"nutmeg":               {},
	"oregano":              {},
	"paprika":              {},
	"parsley":              {},
	"perilla":              {},
	"peruvian pepper":      {},
	"pomegranate seed":     {},
	"poppy seed":           {},
	"radhuni":              {},
	"rose":                 {},
	"rosemary":             {},
	"rue":                  {},
	"saffron":              {},
	"sage":                 {},
	"salt":                 {},
	"sansho":               {},
	"sarsaparilla":         {},
	"sassafras":            {},
	"savory":               {},
	"sesame":               {},
	"shiso":                {},
	"sichuan pepper":       {},
	"sorrel":               {},
	"star anise":           {},
	"sumac":                {},
	"tamarind":             {},
	"tarragon":             {},
	"tasmanian pepper":     {},
	"thai basil":           {},
	"thyme":                {},
	"tonka bean":           {},
	"turmeric":             {},
	"uzazi":                {},
	"vanilla":              {},
	"vietnamese coriander": {},
	"voatsiperifery":       {},
	"wasabi":               {},
	"white mustard":        {},
	"white peppercorn":     {},
	"woodruff":             {},
	"yuzu":                 {},
	"zedoary":              {},
	"zereshk":              {},
	"zest":                 {},
}

var fruitMap = map[string]struct{}{
	"apple":             {},
	"apricot":           {},
	"avocado":           {},
	"banana":            {},
	"bell pepper":       {},
	"bilberry":          {},
	"blackberry":        {},
	"blackcurrant":      {},
	"blood orange":      {},
	"blueberry":         {},
	"boysenberry":       {},
	"breadfruit":        {},
	"canary melon":      {},
	"cantaloupe":        {},
	"cherimoya":         {},
	"cherry":            {},
	"chili pepper":      {},
	"clementine":        {},
	"cloudberry":        {},
	"coconut":           {},
	"cranberry":         {},
	"cucumber":          {},
	"currant":           {},
	"damson":            {},
	"date":              {},
	"dragonfruit":       {},
	"durian":            {},
	"eggplant":          {},
	"elderberry":        {},
	"feijoa":            {},
	"fig":               {},
	"goji berry":        {},
	"gooseberry":        {},
	"grape":             {},
	"grapefruit":        {},
	"guava":             {},
	"honeydew":          {},
	"huckleberry":       {},
	"jackfruit":         {},
	"jambul":            {},
	"jujube":            {},
	"kiwi fruit":        {},
	"kumquat":           {},
	"lemon":             {},
	"lime":              {},
	"loquat":            {},
	"lychee":            {},
	"mandarine":         {},
	"mango":             {},
	"mulberry":          {},
	"nectarine":         {},
	"nut":               {},
	"olive":             {},
	"orange":            {},
	"pamelo":            {},
	"papaya":            {},
	"passionfruit":      {},
	"peach":             {},
	"pear":              {},
	"persimmon":         {},
	"physalis":          {},
	"pineapple":         {},
	"plum":              {},
	"pomegranate":       {},
	"pomelo":            {},
	"purple mangosteen": {},
	"quince":            {},
	"raisin":            {},
	"rambutan":          {},
	"raspberry":         {},
	"redcurrant":        {},
	"rock melon":        {},
	"salal berry":       {},
	"satsuma":           {},
	"star fruit":        {},
	"strawberry":        {},
	"tamarillo":         {},
	"tangerine":         {},
	"ugli fruit":        {},
	"watermelon":        {},
}

var vegetableMap = map[string]struct{}{
	"acorn squash":        {},
	"alfalfa sprout":      {},
	"amaranth":            {},
	"anise":               {},
	"artichoke":           {},
	"arugula":             {},
	"asparagus":           {},
	"aubergine":           {},
	"azuki bean":          {},
	"banana squash":       {},
	"basil":               {},
	"bean sprout":         {},
	"beet":                {},
	"black bean":          {},
	"black-eyed pea":      {},
	"bok choy":            {},
	"borlotti bean":       {},
	"broad beans":         {},
	"broccoflower":        {},
	"broccoli":            {},
	"brussels sprout":     {},
	"butternut squash":    {},
	"cabbage":             {},
	"calabrese":           {},
	"caraway":             {},
	"carrot":              {},
	"cauliflower":         {},
	"cayenne pepper":      {},
	"celeriac":            {},
	"celery":              {},
	"chamomile":           {},
	"chard":               {},
	"chayote":             {},
	"chickpea":            {},
	"chives":              {},
	"cilantro":            {},
	"collard green":       {},
	"corn":                {},
	"corn salad":          {},
	"courgette":           {},
	"cucumber":            {},
	"daikon":              {},
	"delicata":            {},
	"dill":                {},
	"eggplant":            {},
	"endive":              {},
	"fennel":              {},
	"fiddlehead":          {},
	"frisee":              {},
	"garlic":              {},
	"gem squash":          {},
	"ginger":              {},
	"green bean":          {},
	"green pepper":        {},
	"habanero":            {},
	"herbs and spice":     {},
	"horseradish":         {},
	"hubbard squash":      {},
	"jalapeno":            {},
	"jerusalem artichoke": {},
	"jicama":              {},
	"kale":                {},
	"kidney bean":         {},
	"kohlrabi":            {},
	"lavender":            {},
	"leek ":               {},
	"legume":              {},
	"lemon grass":         {},
	"lentils":             {},
	"lettuce":             {},
	"lima bean":           {},
	"mamey":               {},
	"mangetout":           {},
	"marjoram":            {},
	"mung bean":           {},
	"mushroom":            {},
	"mustard green":       {},
	"navy bean":           {},
	"new zealand spinach": {},
	"nopale":              {},
	"okra":                {},
	"onion":               {},
	"oregano":             {},
	"paprika":             {},
	"parsley":             {},
	"parsnip":             {},
	"patty pan":           {},
	"pea":                 {},
	"pinto bean":          {},
	"potato":              {},
	"pumpkin":             {},
	"radicchio":           {},
	"radish":              {},
	"rhubarb":             {},
	"rosemary":            {},
	"runner bean":         {},
	"rutabaga":            {},
	"sage":                {},
	"scallion":            {},
	"shallot":             {},
	"skirret":             {},
	"snap pea":            {},
	"soy bean":            {},
	"spaghetti squash":    {},
	"spinach":             {},
	"squash ":             {},
	"sweet potato":        {},
	"tabasco pepper":      {},
	"taro":                {},
	"tat soi":             {},
	"thyme":               {},
	"topinambur":          {},
	"tubers":              {},
	"turnip":              {},
	"wasabi":              {},
	"water chestnut":      {},
	"watercress":          {},
	"white radish":        {},
	"yam":                 {},
	"zucchini":            {},
}

var corpusIngredients = []string{" can cream of chicken soup
	" nonhydrogenated margarine
	" can cream of chicken soup ",
	" nonhydrogenated margarine ",
	" pomegranate concentrate
	" cream of mushroom soup
	" pomegranate concentrate ",
	" cocktail pumpernickel
	" cream of mushroom soup ",
	" graham cracker crumbs
	" miniature marshmallow
	" raspberry vinaigrette
	" salt and black pepper
	" balsamic vinaigrette
	" blueberry blackberry
	" butterscotch pudding
	" chanterelle mushroom
	" cocktail pumpernickel ",
	" decaffeinated coffee
	" graham cracker crumbs ",
	" heavy whipping cream
	" jagermeister liqueur
	" meringue buttercream
	" miniature marshmallow ",
	" monosodium glutamate
	" raspberry blackberry
	" raspberry vinaigrette ",
	" salt and black pepper ",
	" tamarind concentrate
	" vegetable shortening
	" vinaigrette dressing
	" worcestershire sauce
	" apple cider vinegar
	" assorted decoration
	" balsamic vinaigrette ",
	" benedictine liqueur
	" blueberry blackberry ",
	" butterscotch pudding ",
	" chanterelle mushroom ",
	" chipotle mayonnaise
	" confectioners sugar
	" crystallized ginger
	" cultured buttermilk
	" decaffeinated coffee ",
	" elderflower cordial
	" elderflower liqueur
	" guanciale pancettum
	" heavy whipping cream ",
	" hellmann mayonnaise
	" jagermeister liqueur ",
	" jerusalem artichoke
	" limeade concentrate
	" limnophila aromaticas ",
	" marinated artichoke
	" meringue buttercream ",
	" monosodium glutamate ",
	" orange flower water
	" parmigiano reggiano
	" pepperoncini pepper
	" portabella mushroom
	" portobello mushroom
	" raspberry blackberry ",
	" shallot vinaigrette
	" sweetened cranberry
	" szechuan peppercorn
	" tamarind concentrate ",
	" vegetable shortening ",
	" vietnamese corianders ",
	" vinaigrette dressing ",
	" worcestershire sauce ",
	" amontillado sherry
	" appenzeller cheese
	" apple cider vinegar ",
	" assorted decoration ",
	" assorted vegetable
	" balsamic reduction
	" benedictine liqueur ",
	" bucatini spaghetti
	" buffalo mozzarella
	" buttermilk biscuit
	" butterscotch sauce
	" cabernet sauvignon
	" californium walnut
	" chartreuse liqueur
	" cheesecake filling
	" chicken tenderloin
	" chilled buttermilk
	" chipotle mayonnaise ",
	" chocolate frosting
	" clementine segment
	" confectioners sugar ",
	" corkscrew macaroni
	" crystalized ginger
	" crystallized ginger ",
	" cultured buttermilk ",
	" desiccated coconut
	" elderflower cordial ",
	" elderflower liqueur ",
	" guanciale pancettum ",
	" hellmann mayonnaise ",
	" jerusalem artichoke ",
	" jerusalem artichokes ",
	" limeade concentrate ",
	" limnophila aromatica ",
	" limoncello liqueur
	" luxardo maraschino
	" maraschino liqueur
	" marinated artichoke ",
	" nasturtium blossom
	" new zealand spinachs ",
	" orange flower water ",
	" parmigiano reggiano ",
	" peppermint schnapp
	" pepperoncini pepper ",
	" portabella mushroom ",
	" portobello mushroom ",
	" portuguese sausage
	" precooked polentum
	" pumpernickel bread
	" quality mayonnaise
	" romanesco broccoli
	" rotisserie chicken
	" shallot vinaigrette ",
	" shichimi togarashi
	" sichuan peppercorn
	" sourdough baguette
	" spaghetti linguine
	" spinach fettuccine
	" strawberry gelatin
	" sweetened cranberry ",
	" szechuan peppercorn ",
	" vegetable bouillon
	" vegetarian sausage
	" venison tenderloin
	" vietnamese coriander ",
	" white wine vinegar
	" yellowtail snapper
	" amontillado sherry ",
	" andouille sausage
	" appenzeller cheese ",
	" apple concentrate
	" assorted mushroom
	" assorted vegetable ",
	" balsamic reduction ",
	" barbecued chicken
	" black peppercorns
	" blackberry brandy
	" blanched hazelnut
	" bolivian corianders ",
	" breakfast sausage
	" broccoli floweret
	" bucatini spaghetti ",
	" buffalo mozzarella ",
	" buttermilk biscuit ",
	" butterscotch chip
	" butterscotch sauce ",
	" cabernet sauvignon ",
	" californium walnut ",
	" candy thermometer
	" caramelized onion
	" cellophane noodle
	" champagne vinegar
	" chartreuse liqueur ",
	" cheese tortellini
	" cheesecake filling ",
	" chicken tenderloin ",
	" chilled buttermilk ",
	" chilled champagne
	" chimichurri sauce
	" chinkiang vinegar
	" chocolate biscuit
	" chocolate frosting ",
	" chocolate liqueur
	" chocolate pudding
	" chocolate shaving
	" chocolate soymilk
	" chunky applesauce
	" clementine segment ",
	" cointreau liqueur
	" corkscrew macaroni ",
	" cornstarch sifted
	" cranberry liqueur
	" cream of mushroom
	" crisco shortening
	" crystalized ginger ",
	" cultivated mussel
	" desiccated coconut ",
	" digestive biscuit
	" dijon vinaigrette
	" emmentaler cheese
	" european cucumber
	" fingerling potato
	" flatbread cracker
	" framboise liqueur
	" gingerbread spice
	" gingersnap cookie
	" gorgonzola cheese
	" grains of paradises ",
	" granular fructose
	" granulated garlic
	" green bell pepper
	" hardwood charcoal
	" horseradish sauce
	" hothouse cucumber
	" houttuynia cordatas ",
	" hulled strawberry
	" hungarian paprika
	" imitation vanilla
	" italian seasoning
	" japanese cucumber
	" japanese eggplant
	" jerusalem artichoke ",
	" juice concentrate
	" kefalotyri cheese
	" kewpie mayonnaise
	" lebanese cucumber
	" lemon vinaigrette
	" limoncello liqueur ",
	" lowfat buttermilk
	" lowfat mayonnaise
	" luxardo maraschino ",
	" malaguetum pepper
	" manzanilla sherry
	" maraschino cherry
	" maraschino liqueur ",
	" marshmallow cream
	" marshmallow creme
	" marshmallow fluff
	" mascarpone cheese
	" mccormick paprika
	" mexican chocolate
	" mozzarella cheese
	" nasturtium blossom ",
	" neufchatel cheese
	" new zealand spinach ",
	" nonfat buttermilk
	" nonfat mayonnaise
	" parmesan pecorino
	" peekytoe crabmeat
	" peppermint schnapp ",
	" pickled vegetable
	" pickling cucumber
	" pomegranate juice
	" pomegranate syrup
	" porterhouse steak
	" portuguese sausage ",
	" precooked polentum ",
	" pumpernickel bread ",
	" pumpkin pie spice
	" quality mayonnaise ",
	" raspberry liqueur
	" raspberry vinegar
	" red pepper flakes
	" romanesco broccoli ",
	" rotisserie chicken ",
	" shaken buttermilk
	" shichimi togarashi ",
	" shiitake mushroom
	" shoestring potato
	" shortbread cookie
	" shortcrust pastry
	" sichuan peppercorn ",
	" slender asparagus
	" smoked mozzarella
	" sourdough baguette ",
	" sourdough crouton
	" sourdough starter
	" spaghetti linguine ",
	" spinach fettuccine ",
	" strawberry gelatin ",
	" strawberry yogurt
	" sweetened coconut
	" turkey tenderloin
	" unblanched almond
	" vanilla flavoring
	" vanilla ice cream
	" vegetable bouillon ",
	" vegetarian sausage ",
	" venison tenderloin ",
	" vermicelli noodle
	" watermelon radish
	" wheat orecchiette
	" white horseradish
	" white wine vinegar ",
	" whole wheat flour
	" yellow chartreuse
	" yellowtail snapper ",
	" additional salsa
	" amaretti biscuit
	" amaretto liqueur
	" amaro montenegro
	" andouille sausage ",
	" apple concentrate ",
	" artichoke bottom
	" artichoke hearts
	" assorted filling
	" assorted mushroom ",
	" assorted topping
	" baking chocolate
	" balsamic vinegar
	" barbecued chicken ",
	" beefsteak tomato
	" berry strawberry
	" black peppercorns ",
	" blackberry brandy ",
	" blackberry syrup
	" blanched hazelnut ",
	" bolivian coriander ",
	" bottle champagne
	" breakfast cereal
	" breakfast radish
	" breakfast sausage ",
	" broccoli florets
	" broccoli floweret ",
	" buckwheat noodle
	" butter margarine
	" buttermilk bread
	" butternut squash
	" butterscotch chip ",
	" camembert cheese
	" canadian whiskey
	" candy thermometer ",
	" caramelized onion ",
	" carbonated water
	" cellophane noodle ",
	" champagne vinegar ",
	" cheese tortellini ",
	" cherrystone clam
	" chicken bouillon
	" chicken drumette
	" chihuahua cheese
	" chilled champagne ",
	" chilled prosecco
	" chimichurri sauce ",
	" chinese broccoli
	" chinese eggplant
	" chinkiang vinegar ",
	" chocolate biscuit ",
	" chocolate cookie
	" chocolate liqueur ",
	" chocolate morsel
	" chocolate mousse
	" chocolate pudding ",
	" chocolate shaving ",
	" chocolate soymilk ",
	" chocolate spread
	" chunky applesauce ",
	" clarified butter
	" clementine juice
	" cocchi americano
	" cocktail sausage
	" coconut crystals
	" coconut macaroon
	" cointreau liqueur ",
	" compressed yeast
	" cornflake cereal
	" cornichon pickle
	" cornish game hen
	" cornstarch sifted ",
	" cranberry liqueur ",
	" cream of mushroom ",
	" cremini mushroom
	" crimini mushroom
	" crisco shortening ",
	" cubanelle pepper
	" cultivated mussel ",
	" dandelion flower
	" decorative sugar
	" delicatum squash
	" digestive biscuit ",
	" dijon mayonnaise
	" dijon vinaigrette ",
	" dumpling wrapper
	" east asian peppers ",
	" emmentaler cheese ",
	" empanada wrapper
	" english cucumber
	" espelette pepper
	" european cucumber ",
	" fingerling potato ",
	" flatbread cracker ",
	" framboise liqueur ",
	" german chocolate
	" ginger marmalade
	" gingerbread spice ",
	" gingersnap cookie ",
	" gorgonzola cheese ",
	" gorgonzola dolce
	" grains of paradise ",
	" granular fructose ",
	" granulated garlic ",
	" grapefruit juice
	" green bell pepper ",
	" green chartreuse
	" green peppercorn
	" hardwood charcoal ",
	" hazelnut liqueur
	" homogenized milk
	" honeycrisp apple
	" horseradish sauce ",
	" hot pepper sauce
	" hothouse cucumber ",
	" houttuynia cordata ",
	" hulled strawberry ",
	" hungarian paprika ",
	" imitation vanilla ",
	" instant espresso
	" instant polentum
	" italian dressing
	" italian eggplant
	" italian meatball
	" italian parmesan
	" italian seasoning ",
	" jalapeno chilies
	" japanese cucumber ",
	" japanese eggplant ",
	" japanese pumpkin
	" jarlsberg cheese
	" juice concentrate ",
	" kefalotyri cheese ",
	" kentucky bourbon
	" kewpie mayonnaise ",
	" kielbasa sausage
	" lebanese cucumber ",
	" lemon vinaigrette ",
	" linguica sausage
	" lowfat buttermilk ",
	" lowfat margarine
	" lowfat mayonnaise ",
	" maitake mushroom
	" malaguetum pepper ",
	" manzanilla sherry ",
	" maraschino cherry ",
	" marshmallow cream ",
	" marshmallow creme ",
	" marshmallow fluff ",
	" marshmallow peep
	" mascarpone cheese ",
	" mccormick paprika ",
	" mexican chocolate ",
	" mixed peppercorn
	" mozzarella cheese ",
	" multigrain bread
	" mushroom risotto
	" neufchatel cheese ",
	" nondairy creamer
	" nonfat buttermilk ",
	" nonfat mayonnaise ",
	" orange marmalade
	" panko breadcrumb
	" pareve margarine
	" parmesan crouton
	" parmesan pecorino ",
	" parsley cilantro
	" pecorino toscano
	" peekytoe crabmeat ",
	" pencil asparagus
	" pickled cucumber
	" pickled jalapeno
	" pickled vegetable ",
	" pickling cucumber ",
	" pomegranate juice ",
	" pomegranate seed
	" pomegranate syrup ",
	" porcini mushroom
	" porterhouse steak ",
	" processed cheese
	" prosciutto cotto
	" provolone cheese
	" pumpkin pie spice ",
	" purple mangosteens ",
	" raspberry liqueur ",
	" raspberry sorbet
	" raspberry vinegar ",
	" red pepper flakes ",
	" red wine vinegar
	" reposado tequila
	" roasting chicken
	" roquefort cheese
	" saffron optional
	" shaken buttermilk ",
	" shaved chocolate
	" shiitake mushroom ",
	" shitake mushroom
	" shoestring potato ",
	" shortbread cookie ",
	" shortbread crust
	" shortcrust pastry ",
	" skinned hazelnut
	" slender asparagus ",
	" smoked bratwurst
	" smoked mozzarella ",
	" smoked whitefish
	" solid shortening
	" sourdough crouton ",
	" sourdough starter ",
	" southern comfort
	" spaghetti noodle
	" spaghetti squash
	" spanish mackerel
	" spinach linguine
	" spinach tortilla
	" splash grenadine
	" splenda granular
	" square focaccium
	" strawberry juice
	" strawberry puree
	" strawberry sauce
	" strawberry syrup
	" strawberry yogurt ",
	" sugar substitute
	" sunflower sprout
	" sweetened coconut ",
	" tarragon parsley
	" tarragon vinegar
	" tenderloin roast
	" togarashi pepper
	" turkey pepperoni
	" turkey tenderloin ",
	" unblanched almond ",
	" unbleached flour
	" unprocessed bran
	" valencium orange
	" vanilla flavoring ",
	" vanilla frosting
	" vanilla ice cream ",
	" vegan mayonnaise
	" vegetarian bacon
	" vegetarian chili
	" velveetum cheese
	" vermicelli noodle ",
	" watermelon juice
	" watermelon radish ",
	" wheat breadcrumb
	" wheat fettuccine
	" wheat orecchiette ",
	" white grapefruit
	" white horseradish ",
	" white peppercorn
	" whole wheat flour ",
	" wildflower honey
	" winter vegetable
	" yellow chartreuse ",
	" zucchini blossom
	" additional salsa ",
	" alligator peppers ",
	" amaretti biscuit ",
	" amaretti cookie
	" amaretto liqueur ",
	" amaro montenegro ",
	" american cheese
	" apricot gelatin
	" apricot liqueur
	" armagnac cognac
	" arrowroot flour
	" artichoke bottom ",
	" artichoke heart
	" artichoke hearts ",
	" assorted filling ",
	" assorted topping ",
	" atlantic salmon
	" bacon pancettum
	" baking chocolate ",
	" balsamic vinegar ",
	" beef tenderloin
	" beefsteak tomato ",
	" berry strawberry ",
	" black peppercorns ",
	" blackberry syrup ",
	" blanched almond
	" bloody mary mix
	" blueberry juice
	" bottle champagne ",
	" bourbon whiskey
	" brazilian peppers ",
	" breakfast cereal ",
	" breakfast radish ",
	" brewed espresso
	" broccoli florets ",
	" broiler chicken
	" buckwheat flour
	" buckwheat groat
	" buckwheat honey
	" buckwheat noodle ",
	" burratum cheese
	" butter margarine ",
	" butter optional
	" buttered noodle
	" buttermilk bread ",
	" butternut squash ",
	" butternut squashs ",
	" buttery cracker
	" button mushroom
	" cacao chocolate
	" caesar dressing
	" calamatum olive
	" camembert cheese ",
	" canadian whiskey ",
	" cannellini bean
	" caramel topping
	" carbonated water ",
	" cerignola olive
	" challah brioche
	" champagne grape
	" champagne yeast
	" charentai melon
	" cherry tomatoes
	" cherrystone clam ",
	" chervil parsley
	" chicken bouillon ",
	" chicken breasts
	" chicken carcass
	" chicken drumette ",
	" chicken gizzard
	" chicken sausage
	" chickpea rinsed
	" chicory lettuce
	" chihuahua cheese ",
	" chilled prosecco ",
	" chilled seltzer
	" chinese broccoli ",
	" chinese cabbage
	" chinese eggplant ",
	" chinese mustard
	" chinese sausage
	" chipotle pepper
	" chocolate cookie ",
	" chocolate crust
	" chocolate morsel ",
	" chocolate mousse ",
	" chocolate sauce
	" chocolate spread ",
	" chocolate syrup
	" chocolate wafer
	" chorizo sausage
	" ciabattum bread
	" cilantro leaves
	" cipolline onion
	" cipollini onion
	" clarified butter ",
	" clementine juice ",
	" cocchi americano ",
	" cocktail peanut
	" cocktail sausage ",
	" coconut crystals ",
	" coconut macaroon ",
	" coconut shaving
	" coconut vinegar
	" common mushroom
	" compressed yeast ",
	" cornbread crumb
	" cornflake cereal ",
	" cornflake crumb
	" cornichon pickle ",
	" cornish game hen ",
	" country sausage
	" cranberry jelly
	" cranberry juice
	" cranberry sauce
	" cranberry vodka
	" cream of tartar
	" cremini mushroom ",
	" crimini mushroom ",
	" croissant dough
	" crusty baguette
	" cubanelle pepper ",
	" cultured butter
	" dandelion flower ",
	" dandelion green
	" decorative sugar ",
	" delicatum squash ",
	" desired topping
	" dijon mayonnaise ",
	" dumpling wrapper ",
	" east asian pepper ",
	" elephant garlic
	" emmental cheese
	" empanada wrapper ",
	" enchilada sauce
	" english cucumber ",
	" english mustard
	" espelette pepper ",
	" espresso coffee
	" evaporated milk
	" florida avocado
	" flour tortillas
	" focaccium bread
	" focaccium dough
	" freestone peach
	" french baguette
	" gelatin dessert
	" german chocolate ",
	" ginger marmalade ",
	" globe artichoke
	" golden flaxseed
	" gorgonzola dolce ",
	" graham crackers
	" grapefruit juice ",
	" grapefruit peel
	" greater galangals ",
	" green asparagus
	" green chartreuse ",
	" green coriander
	" green peppercorn ",
	" green peppercorns ",
	" grenadine syrup
	" habanero pepper
	" halloumi cheese
	" hazelnut butter
	" hazelnut liqueur ",
	" hazelnut spread
	" hefeweizen beer
	" homogenized milk ",
	" honeycrisp apple ",
	" hot pepper sauce ",
	" iceberg lettuce
	" iceburg lettuce
	" instant couscou
	" instant espresso ",
	" instant polentum ",
	" instant tapioca
	" israeli couscou
	" italian dressing ",
	" italian eggplant ",
	" italian fontina
	" italian meatball ",
	" italian parmesan ",
	" italian parsley
	" italian sausage
	" jalapeno chilies ",
	" jalapeno pepper
	" japanese pumpkin ",
	" jarlsberg cheese ",
	" kalamatum olive
	" kentucky bourbon ",
	" kielbasa sausage ",
	" kitchen bouquet
	" lamb tenderloin
	" lasagna noodles
	" lavender flower
	" lemon marmalade
	" lime mayonnaise
	" lingonberry jam
	" linguica sausage ",
	" littleneck clam
	" lowfat margarine ",
	" lumpium wrapper
	" macaroni noodle
	" macintosh apple
	" maitake mushroom ",
	" manchego cheese
	" mandarin orange
	" margaritum salt
	" marshmallow peep ",
	" merguez sausage
	" mexican chorizo
	" mexican oregano
	" milk mozzarella
	" miniature bagel
	" mixed peppercorn ",
	" mixed vegetable
	" muenster cheese
	" multigrain bread ",
	" muscovado sugar
	" mushroom risotto ",
	" mzarella cheese
	" nondairy creamer ",
	" nonpareil caper
	" orange marmalade ",
	" oyster mushroom
	" pancettum bacon
	" panko breadcrumb ",
	" pareve margarine ",
	" parmesan cheese
	" parmesan crouton ",
	" parmesan shaved
	" parsley chervil
	" parsley cilantro ",
	" pasteurized egg
	" pattypan squash
	" peach nectarine
	" pecorino cheese
	" pecorino romano
	" pecorino toscano ",
	" pencil asparagus ",
	" peppadew pepper
	" pickled cucumber ",
	" pickled herring
	" pickled jalapeno ",
	" pickled mustard
	" piloncillo cone
	" pimiento pepper
	" pineapple juice
	" pineapple salsa
	" piquillo pepper
	" pistachio pecan
	" polish kielbasa
	" polska kielbasa
	" pomegranate seed ",
	" pomegranate seeds ",
	" porcini mushroom ",
	" pork tenderloin
	" premium tequila
	" pressure cooker
	" processed cheese ",
	" prosciutto cotto ",
	" provolone cheese ",
	" purple mangosteen ",
	" raclette cheese
	" raspberry couli
	" raspberry juice
	" raspberry puree
	" raspberry sauce
	" raspberry sorbet ",
	" raspberry vodka
	" red bell pepper
	" red wine vinegar ",
	" remoulade sauce
	" reposado tequila ",
	" rice vermicelli
	" roasting chicken ",
	" romaine lettuce
	" roquefort cheese ",
	" rosemary branch
	" ruby grapefruit
	" saffron optional ",
	" salt and pepper
	" saltine cracker
	" saskatoon berry
	" sauvignon blanc
	" seltzer chilled
	" shaved chocolate ",
	" shaved parmesan
	" shishito pepper
	" shitake mushroom ",
	" shortbread crust ",
	" shrimp bouillon
	" skinned hazelnut ",
	" smoked bluefish
	" smoked bratwurst ",
	" smoked kielbasa
	" smoked mackerel
	" smoked whitefish ",
	" solid shortening ",
	" sourdough bread
	" southern comfort ",
	" spaetzle noodle
	" spaghetti noodle ",
	" spaghetti squash ",
	" spaghetti squashs ",
	" spanish chorizo
	" spanish mackerel ",
	" spanish paprika
	" sparkling cider
	" sparkling sugar
	" sparkling water
	" sparkling white
	" spinach linguine ",
	" spinach souffle
	" spinach tortilla ",
	" splash grenadine ",
	" splenda granular ",
	" square focaccium ",
	" stewing chicken
	" strawberry juice ",
	" strawberry puree ",
	" strawberry sauce ",
	" strawberry syrup ",
	" sugar substitute ",
	" sundried tomato
	" sunflower sprout ",
	" superfine sugar
	" swordfish steak
	" szechwan pepper
	" taleggio cheese
	" tandoori masala
	" tangerine juice
	" tarragon parsley ",
	" tarragon vinegar ",
	" tasmanian peppers ",
	" tenderloin roast ",
	" thick asparagus
	" togarashi pepper ",
	" tomatillo salsa
	" tomatillo sauce
	" tomato bouillon
	" tomato passatum
	" tomato tapenade
	" turbinado sugar
	" turkey dripping
	" turkey kielbasa
	" turkey pepperoni ",
	" turkey stuffing
	" turkish apricot
	" unbleached flour ",
	" unprocessed bran ",
	" unrefined sugar
	" unripe plantain
	" unseasoned rice
	" valencium orange ",
	" vanilla custard
	" vanilla essence
	" vanilla frosting ",
	" vanilla pudding
	" veal scallopini
	" veal scaloppine
	" vegan chocolate
	" vegan margarine
	" vegan mayonnaise ",
	" vegetable broth
	" vegetable juice
	" vegetable spray
	" vegetable stock
	" vegetarian bacon ",
	" vegetarian chili ",
	" velveetum cheese ",
	" water chestnuts
	" watermelon juice ",
	" wheat breadcrumb ",
	" wheat fettuccine ",
	" wheat spaghetti
	" white asparagus
	" white chocolate
	" white grapefruit ",
	" white peppercorn ",
	" white peppercorns ",
	" wholemeal flour
	" wildflower honey ",
	" winter vegetable ",
	" worchestershire
	" yellow cake mix
	" yellow capsicum
	" yellow cornmeal
	" yellow plantain
	" yellow zucchini
	" zucchini blossom ",
	" zucchini flower
	" zucchini noodle
	" alfalfa sprout
	" alligator meat
	" alligator pepper ",
	" allspice berry
	" almond essence
	" almond extract
	" almond liqueur
	" alphonso mango
	" amaranth flour
	" amarena cherry
	" amaretti cookie ",
	" american cheese ",
	" anaheim pepper
	" anned chickpea
	" apricot brandy
	" apricot gelatin ",
	" apricot halved
	" apricot liqueur ",
	" apricot nectar
	" armagnac cognac ",
	" aromatic gingers ",
	" arrowroot flour ",
	" artichoke heart ",
	" asadero cheese
	" asian eggplant
	" assorted fruit
	" assorted green
	" assorted olive
	" atlantic salmon ",
	" baby artichoke
	" baby asparagus
	" bacon dripping
	" bacon pancettum ",
	" banana liqueur
	" banyul vinegar
	" barbecue sauce
	" barbecued pork
	" barbeque sauce
	" bechamel sauce
	" beef tenderloin ",
	" belgian endife
	" belgian endive
	" black peppercorn ",
	" blackberry jam
	" blanc vermouth
	" blanched almond ",
	" blanco tequila
	" bloody mary mix ",
	" blueberry juice ",
	" boston lettuce
	" bottle ketchup
	" bourbon brandy
	" bourbon whiskey ",
	" boursin cheese
	" braeburn apple
	" braising green
	" bratwurst link
	" brazilian pepper ",
	" bread stuffing
	" brewed espresso ",
	" broccoli crown
	" broiler chicken ",
	" brown flaxseed
	" browned butter
	" brussel sprout
	" brussels sprouts ",
	" brut champagne
	" buckwheat flour ",
	" buckwheat groat ",
	" buckwheat honey ",
	" buckwheat soba
	" budweiser beer
	" bunium persicums ",
	" burratum cheese ",
	" butter lettuce
	" butter optional ",
	" buttered noodle ",
	" buttered toast
	" butternut squash ",
	" buttery cracker ",
	" button mushroom ",
	" cabbage kimchi
	" cacao chocolate ",
	" caesar dressing ",
	" calamatum olive ",
	" canadian bacon
	" candied cherry
	" candied citron
	" candied ginger
	" candied walnut
	" cannellini bean ",
	" caramel topping ",
	" carnaroli rice
	" cayenne pepper
	" cerignola olive ",
	" challah brioche ",
	" champagne grape ",
	" champagne yeast ",
	" charcoal grill
	" charentai melon ",
	" chayote squash
	" cheddar cheese
	" cheese cracker
	" cherry gelatin
	" cherry heering
	" cherry liqueur
	" cherry tomatoes ",
	" chervil parsley ",
	" chestnut cream
	" chestnut flour
	" chestnut honey
	" chestnut puree
	" chicken breast
	" chicken breasts ",
	" chicken carcass ",
	" chicken giblet
	" chicken gizzard ",
	" chicken nugget
	" chicken sausage ",
	" chicken tender
	" chicken thighs
	" chickpea flour
	" chickpea rinsed ",
	" chicory lettuce ",
	" chilled seltzer ",
	" chinese cabbage ",
	" chinese celery
	" chinese mustard ",
	" chinese sausage ",
	" chipotle chile
	" chipotle chili
	" chipotle pepper ",
	" chipotle puree
	" chipotle salsa
	" chipotle sauce
	" chocolate cake
	" chocolate chip
	" chocolate crust ",
	" chocolate curl
	" chocolate kiss
	" chocolate milk
	" chocolate sauce ",
	" chocolate syrup ",
	" chocolate wafer ",
	" chorizo sausage ",
	" ciabattum bread ",
	" ciabattum loaf
	" ciabattum roll
	" cilantro basil
	" cilantro cream
	" cilantro leaves ",
	" cilantro sauce
	" cinnamon cream
	" cinnamon sugar
	" cipolline onion ",
	" cipollini onion ",
	" cocktail frank
	" cocktail olive
	" cocktail onion
	" cocktail peanut ",
	" cocktail sauce
	" coconut butter
	" coconut shaving ",
	" coconut sorbet
	" coconut vinegar ",
	" coconut yogurt
	" coffee liqueur
	" common cracker
	" common mushroom ",
	" condensed milk
	" converted rice
	" coriander seed
	" corn tortillas
	" cornbread crumb ",
	" cornflake crumb ",
	" cornmeal crust
	" cortland apple
	" cottage cheese
	" country sausage ",
	" cranberry bean
	" cranberry jelly ",
	" cranberry juice ",
	" cranberry sauce ",
	" cranberry vodka ",
	" cream of tartar ",
	" crema mexicana
	" creme anglaise
	" creole mustard
	" crispy shallot
	" croissant dough ",
	" crusty baguette ",
	" cucumber salad
	" cultured butter ",
	" dandelion green ",
	" dark chocolate
	" demerara sugar
	" demerara syrup
	" desired topping ",
	" dry white wine
	" dulce de leche
	" dungeness crab
	" edible glitter
	" egg fettuccine
	" egg substitute
	" elbow macaroni
	" elephant garlic ",
	" emmental cheese ",
	" empanada dough
	" enchilada sauce ",
	" english muffin
	" english mustard ",
	" english walnut
	" enoki mushroom
	" enriched flour
	" espresso coffee ",
	" evaporated milk ",
	" fat buttermilk
	" fat mayonnaise
	" fennel sausage
	" fenugreek seed
	" filtered water
	" finishing salt
	" flageolet bean
	" flatiron steak
	" flavorless oil
	" florida avocado ",
	" flour tortilla
	" flour tortillas ",
	" focaccium bread ",
	" focaccium dough ",
	" fontina cheese
	" forbidden rice
	" freestone peach ",
	" french baguette ",
	" french mustard
	" frying chicken
	" fuyu persimmon
	" garbanzo flour
	" garlic crouton
	" garlic granule
	" garlic sausage
	" gelatin dessert ",
	" german mustard
	" gherkin pickle
	" ginger liqueur
	" globe artichoke ",
	" globe eggplant
	" golden flaxseed ",
	" graham cracker
	" graham crackers ",
	" grains of selims ",
	" granola cereal
	" grapefruit peel ",
	" greater galangal ",
	" green asparagus ",
	" green capsicum
	" green cardamom
	" green coriander ",
	" green peppercorn ",
	" green plantain
	" green scallion
	" green zucchini
	" grenadine syrup ",
	" gruyere cheese
	" guajillo chile
	" guajillo chili
	" guinness stout
	" habanero chile
	" habanero chili
	" habanero pepper ",
	" habanero sauce
	" halloumi cheese ",
	" haloumi cheese
	" hamburger meat
	" hamburger roll
	" hardboiled egg
	" havarti cheese
	" hazelnut butter ",
	" hazelnut flour
	" hazelnut spread ",
	" hazelnut syrup
	" hefeweizen beer ",
	" herbed crouton
	" herbs and spices ",
	" hibiscu flower
	" honeydew melon
	" iceberg lettuce ",
	" iceburg lettuce ",
	" imported olive
	" indian bay leafs ",
	" instant coffee
	" instant couscou ",
	" instant tapioca ",
	" israeli couscou ",
	" italian fontina ",
	" italian parsley ",
	" italian sausage ",
	" italian tomato
	" jalapeno chile
	" jalapeno chili
	" jalapeno pepper ",
	" jonagold apple
	" kabocha squash
	" kalamatum olive ",
	" kasseri cheese
	" kitchen bouquet ",
	" lamb tenderloin ",
	" lasagna noodle
	" lasagna noodles ",
	" lavender flower ",
	" lavender honey
	" lemon marmalade ",
	" lesser galangals ",
	" lettuce tomato
	" lime mayonnaise ",
	" lingonberry jam ",
	" littleneck clam ",
	" loaf ciabattum
	" lumpium wrapper ",
	" luxardo cherry
	" macadamium nut
	" macaroni noodle ",
	" macintosh apple ",
	" manchego cheese ",
	" mandarin juice
	" mandarin orange ",
	" mandarin vodka
	" marcona almond
	" margaritum mix
	" margaritum salt ",
	" marinara sauce
	" marshmallow cr
	" mcintosh apple
	" merguez sausage ",
	" mexican chorizo ",
	" mexican oregano ",
	" mexican tomato
	" mezzi rigatoni
	" milk chocolate
	" milk mozzarella ",
	" miniature bagel ",
	" minute tapioca
	" mixed mushroom
	" mixed vegetable ",
	" morel mushroom
	" muenster cheese ",
	" muscovado sugar ",
	" mushroom broth
	" mushroom sauce
	" mzarella cheese ",
	" natural almond
	" natural yogurt
	" nonpareil caper ",
	" nonstick spray
	" oatmeal cookie
	" olive tapenade
	" onion scallion
	" orange blossom
	" orange curacao
	" orange gelatin
	" orange liqueur
	" oyster cracker
	" oyster mushroom ",
	" oyster shucked
	" package tempeh
	" pancake batter
	" pancettum bacon ",
	" parboiled rice
	" parmesan cheese ",
	" parmesan shaved ",
	" parsley chervil ",
	" parsley flakes
	" parsley leaves
	" pasteurized egg ",
	" pattypan squash ",
	" peach nectarine ",
	" peanut brittle
	" pearled barley
	" pecorino cheese ",
	" pecorino romano ",
	" peppadew pepper ",
	" pepper vinegar
	" peppered bacon
	" peruvian peppers ",
	" pickled garlic
	" pickled ginger
	" pickled herring ",
	" pickled mustard ",
	" pickled pepper
	" pickling spice
	" piloncillo cone ",
	" pimento pepper
	" pimiento pepper ",
	" pineapple juice ",
	" pineapple salsa ",
	" piquillo pepper ",
	" pistachio pecan ",
	" poached salmon
	" poblano pepper
	" polish kielbasa ",
	" polish sausage
	" polska kielbasa ",
	" pomegranate seed ",
	" popcorn kernel
	" popped popcorn
	" pork tenderloin ",
	" portobello cap
	" potato gnocchi
	" powdered sugar
	" premium tequila ",
	" pressure cooker ",
	" prosciutto ham
	" pumpkin butter
	" purified water
	" purple cabbage
	" raclette cheese ",
	" rainbow quinoa
	" raisin currant
	" ranch dressing
	" raspberry couli ",
	" raspberry juice ",
	" raspberry puree ",
	" raspberry sauce ",
	" raspberry vodka ",
	" realemon juice
	" red bell pepper ",
	" remoulade sauce ",
	" rice vermicelli ",
	" ricotta cheese
	" robiola cheese
	" romaine lettuce ",
	" rosemary branch ",
	" rosemary thyme
	" rubber spatula
	" ruby grapefruit ",
	" saffron strand
	" saffron thread
	" salad dressing
	" salt and pepper ",
	" saltine cracker ",
	" sandwich bread
	" saskatoon berry ",
	" satsuma orange
	" sausage casing
	" sauvignon blanc ",
	" scallion green
	" scallion white
	" scotch whiskey
	" seasoning salt
	" seltzer chilled ",
	" semolina flour
	" serrano pepper
	" seville orange
	" shallot halved
	" shaved parmesan ",
	" sherry vinegar
	" shishito pepper ",
	" shrimp bouillon ",
	" shucked oyster
	" sichuan pepper
	" silver tequila
	" simmered pinto
	" smoked bluefish ",
	" smoked kielbasa ",
	" smoked mackerel ",
	" smoked paprika
	" smoked sausage
	" sockeye salmon
	" sourdough bread ",
	" sourdough loaf
	" sourdough roll
	" spaetzle noodle ",
	" spaghetti squash ",
	" spanish brandy
	" spanish chorizo ",
	" spanish paprika ",
	" spanish peanut
	" sparkling cider ",
	" sparkling sugar ",
	" sparkling water ",
	" sparkling white ",
	" sparkling wine
	" spinach souffle ",
	" split chickpea
	" squash blossom
	" sriracha sauce
	" stale baguette
	" starchy potato
	" stewing chicken ",
	" stilton cheese
	" stirred tahini
	" strawberry jam
	" submarine roll
	" sugar snap pea
	" sultana raisin
	" sundried tomato ",
	" sunflower seed
	" superfine sugar ",
	" sweet potatoes
	" swordfish steak ",
	" szechwan pepper ",
	" taleggio cheese ",
	" tamarind juice
	" tandoori masala ",
	" tangerine juice ",
	" tangerine peel
	" tapioca starch
	" tasmanian pepper ",
	" tender lettuce
	" tequila blanco
	" teriyaki sauce
	" thick asparagus ",
	" thyme rosemary
	" tomatillo salsa ",
	" tomatillo sauce ",
	" tomato bouillon ",
	" tomato chutney
	" tomato ketchup
	" tomato passatum ",
	" tomato tapenade ",
	" tonkatsu sauce
	" tortilla flour
	" tropical fruit
	" truffle butter
	" turbinado sugar ",
	" turkey carcass
	" turkey dripping ",
	" turkey kielbasa ",
	" turkey sausage
	" turkey stuffing ",
	" turkish apricot ",
	" tzatziki sauce
	" unrefined sugar ",
	" unripe plantain ",
	" unseasoned rice ",
	" unsmoked bacon
	" vanilla custard ",
	" vanilla essence ",
	" vanilla pudding ",
	" vanilla yogurt
	" veal scallopini ",
	" veal scaloppine ",
	" vegan chocolate ",
	" vegan margarine ",
	" vegetable broth ",
	" vegetable juice ",
	" vegetable soup
	" vegetable spray ",
	" vegetable stock ",
	" vidalium onion
	" vienna sausage
	" wakame seaweed
	" water chestnut
	" water chestnuts ",
	" wheat baguette
	" wheat farfalle
	" wheat macaroni
	" wheat rigatoni
	" wheat spaghetti ",
	" wheat tortilla
	" whipping cream
	" white asparagus ",
	" white chocolate ",
	" white cornmeal
	" white mushroom
	" white peppercorn ",
	" white vermouth
	" wholemeal flour ",
	" wild asparagus
	" wild blueberry
	" wonton wrapper
	" worcestershire
	" worchestershire ",
	" yellow cake mix ",
	" yellow capsicum ",
	" yellow cornmeal ",
	" yellow mustard
	" yellow plantain ",
	" yellow zucchini ",
	" yellowfin tuna
	" young pecorino
	" zucchini flower ",
	" zucchini noodle ",
	" absolut vodka
	" acacium honey
	" albacore tuna
	" aleppo pepper
	" alfalfa sprout ",
	" alfalfa sprouts ",
	" alfredo sauce
	" alligator meat ",
	" allspice berry ",
	" allspice dram
	" almond butter
	" almond essence ",
	" almond extract ",
	" almond liqueur ",
	" almond sliver
	" alphonso mango ",
	" amaranth flour ",
	" amaranth seed
	" amarena cherry ",
	" anaheim chile
	" anaheim chili
	" anaheim pepper ",
	" anchovy filet
	" anise liqueur
	" anned chickpea ",
	" apple chutney
	" apple schnapp
	" apple vinegar
	" apricot brandy ",
	" apricot halved ",
	" apricot nectar ",
	" arame seaweed
	" aromatic ginger ",
	" arugula pesto
	" arugula salad
	" asadero cheese ",
	" asiago cheese
	" asian eggplant ",
	" assorted fruit ",
	" assorted green ",
	" assorted herb
	" assorted olive ",
	" baby artichoke ",
	" baby asparagus ",
	" baby broccoli
	" baby cucumber
	" baby eggplant
	" baby zucchini
	" bacardi limon
	" bacon dripping ",
	" bag cranberry
	" baking potato
	" baking powder
	" bamboo skewer
	" banana liqueur ",
	" banana pepper
	" banyul vinegar ",
	" barbecue sauce ",
	" barbecued pork ",
	" barbeque sauce ",
	" basil parsley
	" bechamel sauce ",
	" beef bouillon
	" beef consomme
	" beef dripping
	" beefeater gin
	" belgian endife ",
	" belgian endive ",
	" beluga lentil
	" black cardamoms ",
	" black-eyed peas ",
	" blackberry jam ",
	" blanc vermouth ",
	" blanco tequila ",
	" blood sausage
	" blueberry jam
	" borlotti bean
	" boston lettuce ",
	" bottle aperol
	" bottle ketchup ",
	" bouquet garni
	" bourbon brandy ",
	" bourbon sauce
	" boursin cheese ",
	" box raspberry
	" braeburn apple ",
	" braising green ",
	" bratwurst link ",
	" bread cracker
	" bread pudding
	" bread stuffing ",
	" brewed coffee
	" broccoli crown ",
	" broccoli rabe
	" broccoli slaw
	" broken walnut
	" brown flaxseed ",
	" brown mustard
	" browned butter ",
	" brussel sprout ",
	" brussels sprout ",
	" brut champagne ",
	" buckwheat soba ",
	" budweiser beer ",
	" buffalo sauce
	" bunium persicum ",
	" burgundy wine
	" butter cookie
	" butter lettuce ",
	" buttered toast ",
	" cabbage kimchi ",
	" canadian bacon ",
	" candied cherry ",
	" candied citron ",
	" candied fruit
	" candied ginger ",
	" candied pecan
	" candied walnut ",
	" candy coating
	" caramel sauce
	" caramel syrup
	" cardamom seed
	" carnaroli rice ",
	" cashew butter
	" cassi liqueur
	" cayenne pepper ",
	" cayenne peppers ",
	" challah bread
	" charcoal grill ",
	" chayote squash ",
	" cheddar cheese ",
	" cheese cracker ",
	" cheese spread
	" cherry brandy
	" cherry gelatin ",
	" cherry heering ",
	" cherry liqueur ",
	" cherry pepper
	" cherry tomato
	" chestnut cream ",
	" chestnut flour ",
	" chestnut honey ",
	" chestnut puree ",
	" chicken breast ",
	" chicken broth
	" chicken giblet ",
	" chicken liver
	" chicken nugget ",
	" chicken salad
	" chicken stock
	" chicken strip
	" chicken tender ",
	" chicken thigh
	" chicken thighs ",
	" chickpea flour ",
	" chilled water
	" chinese celery ",
	" chinese chili
	" chipotle chile ",
	" chipotle chili ",
	" chipotle puree ",
	" chipotle salsa ",
	" chipotle sauce ",
	" chive blossom
	" chocolate bar
	" chocolate cake ",
	" chocolate chip ",
	" chocolate curl ",
	" chocolate egg
	" chocolate kiss ",
	" chocolate milk ",
	" ciabattum loaf ",
	" ciabattum roll ",
	" cider vinegar
	" cilantro basil ",
	" cilantro cream ",
	" cilantro mint
	" cilantro sauce ",
	" cinnamon bark
	" cinnamon cream ",
	" cinnamon sugar ",
	" clam scrubbed
	" clamato juice
	" clotted cream
	" cocktail frank ",
	" cocktail olive ",
	" cocktail onion ",
	" cocktail sauce ",
	" coconut butter ",
	" coconut cream
	" coconut flesh
	" coconut flour
	" coconut sorbet ",
	" coconut sugar
	" coconut water
	" coconut yogurt ",
	" coffee filter
	" coffee liqueur ",
	" collard green
	" colored sugar
	" common cracker ",
	" concord grape
	" condensed milk ",
	" converted rice ",
	" cooking spray
	" coriander leafs ",
	" coriander seed ",
	" coriander seeds ",
	" corn tortilla
	" corn tortillas ",
	" cornmeal crust ",
	" cortland apple ",
	" cotija cheese
	" cottage cheese ",
	" country bread
	" cracker crumb
	" cracker toast
	" cranberry bean ",
	" cranberry jam
	" crawfish tail
	" crayfish tail
	" cream coconut
	" crema mexicana ",
	" creme anglaise ",
	" creme fraiche
	" creole mustard ",
	" crispy shallot ",
	" cucumber salad ",
	" cured chorizo
	" curly parsley
	" curly spinach
	" currant jelly
	" daikon radish
	" dark chocolate ",
	" demerara sugar ",
	" demerara syrup ",
	" dessert apple
	" dijon mustard
	" dinosaur kale
	" dry white wine ",
	" dulce de leche ",
	" dungeness crab ",
	" edible flower
	" edible glitter ",
	" egg fettuccine ",
	" egg substitute ",
	" elbow macaroni ",
	" empanada dough ",
	" english muffin ",
	" english walnut ",
	" enoki mushroom ",
	" enriched flour ",
	" espresso bean
	" farmer cheese
	" fat buttermilk ",
	" fat mayonnaise ",
	" favorite beer
	" fennel pollen
	" fennel sausage ",
	" fenugreek seed ",
	" fernet branca
	" filtered water ",
	" finishing salt ",
	" flageolet bean ",
	" flatiron steak ",
	" flavorless oil ",
	" flaxseed meal
	" flour tortilla ",
	" floury potato
	" fontina cheese ",
	" food coloring
	" forbidden rice ",
	" french lentil
	" french mustard ",
	" fried chicken
	" fried shallot
	" frizzled leek
	" fromage blanc
	" fruit chutney
	" fruit compote
	" fryer chicken
	" frying chicken ",
	" fuyu persimmon ",
	" garbanzo bean
	" garbanzo flour ",
	" garlic confit
	" garlic crouton ",
	" garlic granule ",
	" garlic powder
	" garlic sausage ",
	" garlic sprout
	" gelatin sheet
	" german mustard ",
	" gherkin pickle ",
	" ginger cookie
	" ginger liqueur ",
	" globe eggplant ",
	" glucose syrup
	" glutinou rice
	" golden raisin
	" graham cracker ",
	" grains of selim ",
	" grand marnier
	" granola cereal ",
	" grapeseed oil
	" greek oregano
	" green cabbage
	" green capsicum ",
	" green cardamom ",
	" green chilies
	" green peppers
	" green plantain ",
	" green scallion ",
	" green zucchini ",
	" gruyere cheese ",
	" guajillo chile ",
	" guajillo chili ",
	" guinness stout ",
	" gyoza wrapper
	" habanero chile ",
	" habanero chili ",
	" habanero sauce ",
	" half and half
	" halibut steak
	" haloumi cheese ",
	" hamburger bun
	" hamburger meat ",
	" hamburger roll ",
	" hardboiled egg ",
	" harissa sauce
	" havarti cheese ",
	" hazelnut flour ",
	" hazelnut meal
	" hazelnut skin
	" hazelnut syrup ",
	" heart of palm
	" heart romaine
	" herbed crouton ",
	" herbs and spice ",
	" hibiscu flower ",
	" honey mustard
	" honeydew melon ",
	" hubbard squashs ",
	" imported olive ",
	" inch baguette
	" inch tortilla
	" indian bay leaf ",
	" instant coffee ",
	" instant flour
	" instant yeast
	" irish whiskey
	" italian basil
	" italian bread
	" italian tomato ",
	" jalapeno chile ",
	" jalapeno chili ",
	" jello gelatin
	" jonagold apple ",
	" juniper berry
	" kabocha squash ",
	" kasseri cheese ",
	" kohlrabi bulb
	" kombu seaweed
	" lacinato kale
	" lasagna noodle ",
	" lasagna sheet
	" lavender honey ",
	" lemon gelatin
	" lesser galangal ",
	" lettuce green
	" lettuce tomato ",
	" linguine fini
	" loaf ciabattum ",
	" long eggplant
	" lowfat yogurt
	" luxardo cherry ",
	" macadamium nut ",
	" maine lobster
	" mandarin juice ",
	" mandarin vodka ",
	" mango chutney
	" manioc starch
	" marcona almond ",
	" margaritum mix ",
	" marinara sauce ",
	" marshmallow cr ",
	" mcintosh apple ",
	" mesclun green
	" metaxa brandy
	" mexican crema
	" mexican lager
	" mexican tomato ",
	" mezzi rigatoni ",
	" milk chocolate ",
	" mineral water
	" minute tapioca ",
	" mixed lettuce
	" mixed mushroom ",
	" monterey jack
	" morel mushroom ",
	" muffin batter
	" mushroom broth ",
	" mushroom sauce ",
	" mushroom soup
	" mustard green
	" natural almond ",
	" natural yogurt ",
	" nicoise olive
	" nigella sativas ",
	" nonfat yogurt
	" nonstick spray ",
	" northern bean
	" oatmeal cookie ",
	" oaxaca cheese
	" olive tapenade ",
	" onion scallion ",
	" orange blossom ",
	" orange curacao ",
	" orange gelatin ",
	" orange lentil
	" orange liqueur ",
	" orange pepper
	" orange roughy
	" orange sorbet
	" oyster cracker ",
	" oyster shucked ",
	" package tempeh ",
	" pancake batter ",
	" paneer cheese
	" panela cheese
	" parboiled rice ",
	" parsley basil
	" parsley flakes ",
	" parsley leaves ",
	" parsley pesto
	" parsley thyme
	" pasilla chile
	" passion fruit
	" peach gelatin
	" peach liqueur
	" peach schnapp
	" peanut brittle ",
	" peanut butter
	" pearl couscou
	" pearl tapioca
	" pearled barley ",
	" peasant bread
	" penne fusilli
	" pepper vinegar ",
	" peppered bacon ",
	" peruvian pepper ",
	" phyllo pastry
	" picante salsa
	" picante sauce
	" pickle relish
	" pickled garlic ",
	" pickled ginger ",
	" pickled onion
	" pickled pepper ",
	" pickling lime
	" pickling salt
	" pickling spice ",
	" pimento pepper ",
	" pineapple rum
	" pistachio nut
	" pistachio oil
	" plum tomatoes
	" poached salmon ",
	" poblano chile
	" poblano chili
	" poblano pepper ",
	" poire william
	" polish sausage ",
	" popcorn kernel ",
	" popped popcorn ",
	" pork sparerib
	" portobello cap ",
	" potato gnocchi ",
	" potato starch
	" powdered sugar ",
	" premium vodka
	" prosciutto ham ",
	" puffed millet
	" pullman bread
	" pumpkin butter ",
	" pumpkin flesh
	" pumpkin puree
	" pureed banana
	" pureed tomato
	" purified water ",
	" purple cabbage ",
	" purple potato
	" quick oatmeal
	" radish sprout
	" rainbow chard
	" rainbow quinoa ",
	" rainbow trout
	" raisin currant ",
	" ranch dressing ",
	" raspberry jam
	" realemon juice ",
	" rendered duck
	" rhubarb syrup
	" ricotta cheese ",
	" riesling wine
	" roast chicken
	" robiola cheese ",
	" rolled barley
	" romaine heart
	" romano cheese
	" romesco sauce
	" rosemary thyme ",
	" rubber spatula ",
	" russet potato
	" safflower oil
	" saffron strand ",
	" saffron thread ",
	" salad dressing ",
	" salsa picante
	" sanding sugar
	" sandwich bread ",
	" sandwich roll
	" satsuma orange ",
	" sausage casing ",
	" savoy cabbage
	" scallion green ",
	" scallion white ",
	" scotch bonnet
	" scotch whiskey ",
	" scotch whisky
	" scrambled egg
	" seasoning salt ",
	" seltzer water
	" semolina flour ",
	" serrano chile
	" serrano chili
	" serrano pepper ",
	" sesame tahini
	" seville orange ",
	" shallot halved ",
	" shaoxing wine
	" sharp cheddar
	" shelling bean
	" sherry vinegar ",
	" shucked oyster ",
	" sichuan pepper ",
	" sichuan peppers ",
	" silver tequila ",
	" simmered pinto ",
	" sirloin steak
	" sirloin strip
	" smoked almond
	" smoked paprika ",
	" smoked salmon
	" smoked sausage ",
	" smoked turkey
	" snack cracker
	" sockeye salmon ",
	" sorghum flour
	" sorghum syrup
	" sourdough loaf ",
	" sourdough roll ",
	" soy margarine
	" spanish brandy ",
	" spanish olive
	" spanish onion
	" spanish peanut ",
	" sparkling wine ",
	" spicy mustard
	" spicy sausage
	" spiny lobster
	" split chickpea ",
	" spring garlic
	" squash blossom ",
	" squeeze lemon
	" sriracha sauce ",
	" stale baguette ",
	" starchy potato ",
	" stewed tomato
	" stilton cheese ",
	" stirred tahini ",
	" strawberry jam ",
	" streaky bacon
	" string cheese
	" submarine roll ",
	" sugar pumpkin
	" sugar snap pea ",
	" sultana raisin ",
	" sunflower oil
	" sunflower seed ",
	" sweet potatoes ",
	" tabasco peppers ",
	" tabasco sauce
	" tamarind juice ",
	" tamarind pulp
	" tangerine peel ",
	" tapioca flour
	" tapioca pearl
	" tapioca starch ",
	" tender lettuce ",
	" tequila blanco ",
	" teriyaki sauce ",
	" thyme oregano
	" thyme parsley
	" thyme rosemary ",
	" tomato chutney ",
	" tomato ketchup ",
	" tomato relish
	" tonkatsu sauce ",
	" tortilla chip
	" tortilla flour ",
	" tropical fruit ",
	" truffle butter ",
	" turkey carcass ",
	" turkey sausage ",
	" tzatziki sauce ",
	" unsmoked bacon ",
	" vanilla sauce
	" vanilla sugar
	" vanilla syrup
	" vanilla vodka
	" vanilla wafer
	" vanilla yogurt ",
	" vegan chicken
	" vegetable oil
	" vegetable soup ",
	" veggie burger
	" venison steak
	" vidalium onion ",
	" vienna sausage ",
	" vinegar sauce
	" virginium ham
	" voatsiperiferys ",
	" vodka chilled
	" wakame seaweed ",
	" walnut butter
	" water chestnut ",
	" water chestnuts ",
	" water cracker
	" wheat baguette ",
	" wheat couscou
	" wheat cracker
	" wheat farfalle ",
	" wheat macaroni ",
	" wheat rigatoni ",
	" wheat tortilla ",
	" whipped cream
	" whipping cream ",
	" white cabbage
	" white cheddar
	" white cornmeal ",
	" white creamer
	" white mushroom ",
	" white tequila
	" white truffle
	" white vermouth ",
	" white vinegar
	" wild asparagus ",
	" wild blueberry ",
	" wild mushroom
	" winter squash
	" wonton wrapper ",
	" wooden skewer
	" worcestershire ",
	" yellow hominy
	" yellow lentil
	" yellow mustard ",
	" yellow pepper
	" yellow potato
	" yellow squash
	" yellow tomato
	" yellowfin tuna ",
	" young pecorino ",
	" absolut vodka ",
	" acacium honey ",
	" achiote seed
	" acorn squash
	" agave nectar
	" aged cheddar
	" albacore tuna ",
	" aleppo chili
	" aleppo pepper ",
	" alfalfa sprout ",
	" alfredo sauce ",
	" allspice dram ",
	" almond butter ",
	" almond flour
	" almond sliver ",
	" almond syrup
	" amaranth seed ",
	" anaheim chile ",
	" anaheim chili ",
	" anchovy filet ",
	" anise liqueur ",
	" annatto seed
	" apple brandy
	" apple butter
	" apple chutney ",
	" apple schnapp ",
	" apple vinegar ",
	" arame seaweed ",
	" arborio rice
	" arugula pesto ",
	" arugula salad ",
	" asiago cheese ",
	" assorted herb ",
	" assorted nut
	" averna amaro
	" baby arugula
	" baby broccoli ",
	" baby cucumber ",
	" baby eggplant ",
	" baby lettuce
	" baby spinach
	" baby zucchini ",
	" bacardi limon ",
	" bacon grease
	" bag cranberry ",
	" baked potato
	" bakery bread
	" baking apple
	" baking potato ",
	" baking powder ",
	" baking spray
	" bamboo skewer ",
	" banana bread
	" banana pepper ",
	" banana puree
	" banana squashs ",
	" barbecue rub
	" barley flour
	" basil leaves
	" basil parsley ",
	" basmati rice
	" beef bouillon ",
	" beef brisket
	" beef consomm
	" beef consomme ",
	" beef dripping ",
	" beef sirloin
	" beefeater gin ",
	" beluga lentil ",
	" beurre blanc
	" bibb lettuce
	" black cardamom ",
	" black mustards ",
	" black olives
	" black pepper
	" black-eyed pea ",
	" blood orange
	" blood sausage ",
	" blue curacao
	" blueberry jam ",
	" bone removed
	" borlotti bean ",
	" borlotti beans ",
	" bottle aperol ",
	" bouquet garni ",
	" bourbon sauce ",
	" box raspberry ",
	" bread cracker ",
	" bread crumbs
	" bread pudding ",
	" brewed coffee ",
	" brioche roll
	" broccoli rabe ",
	" broccoli slaw ",
	" broken pecan
	" broken walnut ",
	" brown butter
	" brown lentil
	" brown mustard ",
	" brown mustards ",
	" brown raisin
	" buffalo sauce ",
	" bulb shallot
	" bulgar wheat
	" bulgur wheat
	" bulk sausage
	" burgundy wine ",
	" butter cookie ",
	" butter pecan
	" butter sauce
	" butterscotch
	" calrose rice
	" can chickpea
	" candied fruit ",
	" candied pecan ",
	" candied peel
	" candy coating ",
	" canning salt
	" caramel sauce ",
	" caramel syrup ",
	" caraway seed
	" cardamom pod
	" cardamom seed ",
	" carrot juice
	" cashew butter ",
	" cashew cream
	" cassi liqueur ",
	" caster sugar
	" cayenne pepper ",
	" celery heart
	" celery powders ",
	" chaat masala
	" challah bread ",
	" cheese sauce
	" cheese spread ",
	" cherry brandy ",
	" cherry juice
	" cherry pepper ",
	" cherry syrup
	" cherry tomato ",
	" cherry vodka
	" chicken bone
	" chicken broth ",
	" chicken liver ",
	" chicken meat
	" chicken salad ",
	" chicken skin
	" chicken soup
	" chicken stock ",
	" chicken strip ",
	" chicken thigh ",
	" chicken wing
	" chile pepper
	" chili pepper
	" chili powder
	" chilled lard
	" chilled water ",
	" chinese chili ",
	" chinese wine
	" chipotle rub
	" chive blossom ",
	" chocolate bar ",
	" chocolate egg ",
	" chunky salsa
	" cider vinegar ",
	" cilantro mint ",
	" cinnamon bark ",
	" clam chowder
	" clam scrubbed ",
	" clamato juice ",
	" clotted cream ",
	" clover honey
	" cocoa butter
	" coconut chip
	" coconut cream ",
	" coconut flesh ",
	" coconut flour ",
	" coconut meat
	" coconut milk
	" coconut sugar ",
	" coconut water ",
	" coffee filter ",
	" colby cheese
	" coleslaw mix
	" collard green ",
	" collard greens ",
	" colored sugar ",
	" comte cheese
	" concord grape ",
	" cookie crumb
	" cooking spray ",
	" coriander leaf ",
	" coriander seed ",
	" corn kernels
	" corn shucked
	" corn tortilla ",
	" cotija cheese ",
	" country bread ",
	" country loaf
	" cracker crumb ",
	" cracker toast ",
	" cranberry jam ",
	" crawfish tail ",
	" crayfish tail ",
	" cream cheese
	" cream coconut ",
	" cream sherry
	" creamed corn
	" creme fraiche ",
	" creole spice
	" crusty bread
	" cured chorizo ",
	" curly endive
	" curly parsley ",
	" curly spinach ",
	" currant jelly ",
	" curry powder
	" daikon radish ",
	" demerara rum
	" dessert apple ",
	" dijon mustard ",
	" dill parsley
	" dinosaur kale ",
	" double cream
	" dry red wine
	" edamame bean
	" edible flower ",
	" espresso bean ",
	" farmer cheese ",
	" favorite beer ",
	" fennel frond
	" fennel pollen ",
	" fernet branca ",
	" fetum cheese
	" filet mignon
	" flat anchovy
	" flat parsley
	" flaxseed meal ",
	" flaxseed oil
	" floral honey
	" floury potato ",
	" food coloring ",
	" french bread
	" french lentil ",
	" french toast
	" fresno chile
	" fried chicken ",
	" fried potato
	" fried shallot ",
	" frizzled leek ",
	" fromage blanc ",
	" fruit chutney ",
	" fruit compote ",
	" fryer chicken ",
	" gaetum olive
	" garam masala
	" garbanzo bean ",
	" garlic broth
	" garlic chive
	" garlic chivess ",
	" garlic confit ",
	" garlic hummu
	" garlic juice
	" garlic powder ",
	" garlic scape
	" garlic sprout ",
	" gefilte fish
	" gelatin sheet ",
	" giblet broth
	" gigante bean
	" ginger cookie ",
	" ginger juice
	" ginger syrup
	" globe tomato
	" glucose syrup ",
	" gluten flour
	" glutinou rice ",
	" gold tequila
	" golden raisin ",
	" golden syrup
	" gouda cheese
	" grana padano
	" grand marnier ",
	" grape tomato
	" grapeseed oil ",
	" greek oregano ",
	" greek yogurt
	" green banana
	" green cabbage ",
	" green cherry
	" green chilies ",
	" green garlic
	" green lentil
	" green onions
	" green papaya
	" green pepper
	" green peppers ",
	" green tomato
	" guava nectar
	" gyoza wrapper ",
	" half and half ",
	" halibut steak ",
	" halved lemon
	" hamburger bun ",
	" hanger steak
	" haricot vert
	" harissa sauce ",
	" hazelnut meal ",
	" hazelnut oil
	" hazelnut skin ",
	" heart of palm ",
	" heart romaine ",
	" hearty bread
	" hearty green
	" herb bouquet
	" herb mixture
	" hickory chip
	" hoisin sauce
	" honey mustard ",
	" hubbard squash ",
	" inch baguette ",
	" inch tortilla ",
	" instant flour ",
	" instant grit
	" instant milk
	" instant yeast ",
	" iodized salt
	" irish whiskey ",
	" italian basil ",
	" italian bread ",
	" italian herb
	" italian plum
	" italian tuna
	" jamaican rum
	" jarred salsa
	" jasmine rice
	" jello gelatin ",
	" jigger vodka
	" juice orange
	" jumbo shrimp
	" juniper berry ",
	" juniper berrys ",
	" kidney beans
	" kimchi juice
	" kirschwasser
	" kohlrabi bulb ",
	" kombu seaweed ",
	" lacinato kale ",
	" lasagna sheet ",
	" laughing cow
	" lavender bud
	" lemon gelatin ",
	" lemon halved
	" lemon pepper
	" lemon verbenas ",
	" lemon yogurt
	" lettuce green ",
	" lillet blanc
	" linguine fini ",
	" link sausage
	" litsea cubebas ",
	" little water
	" live lobster
	" loaf brioche
	" loaf challah
	" lobster meat
	" lobster tail
	" london broil
	" long eggplant ",
	" lowfat yogurt ",
	" lychee juice
	" madeira wine
	" maine lobster ",
	" malt vinegar
	" mango chutney ",
	" mango sorbet
	" manioc starch ",
	" marsala wine
	" marshmallows
	" matzo farfel
	" meat filling
	" meatloaf mix
	" medjool date
	" mesclun green ",
	" metal skewer
	" metaxa brandy ",
	" mexican beer
	" mexican crema ",
	" mexican lager ",
	" mexican rice
	" millet flour
	" mineral water ",
	" mint garnish
	" mint parsley
	" miracle whip
	" mixed lettuce ",
	" mixed tomato
	" monterey jack ",
	" muffin batter ",
	" mushroom cap
	" mushroom soup ",
	" mustard green ",
	" mustard greens ",
	" mustard seed
	" napa cabbage
	" navel orange
	" nicoise olive ",
	" nigella sativa ",
	" nigella seed
	" nonfat yogurt ",
	" nori seaweed
	" northern bean ",
	" oaxaca cheese ",
	" onion powder
	" orange juice
	" orange lentil ",
	" orange pepper ",
	" orange roughy ",
	" orange sauce
	" orange sorbet ",
	" orange wheel
	" orgeat syrup
	" oyster sauce
	" oz spaghetti
	" pan dripping
	" panch phoron
	" paneer cheese ",
	" panela cheese ",
	" parsley basil ",
	" parsley dill
	" parsley mint
	" parsley pesto ",
	" parsley thyme ",
	" pasilla chile ",
	" passion fruit ",
	" pastry brush
	" pastry cream
	" pastry dough
	" pastry flour
	" peach brandy
	" peach gelatin ",
	" peach liqueur ",
	" peach nectar
	" peach schnapp ",
	" peach yogurt
	" peanut butter ",
	" peanut sauce
	" pear liqueur
	" pearl barley
	" pearl couscou ",
	" pearl tapioca ",
	" peasant bread ",
	" pecan walnut
	" penne fusilli ",
	" penne rigate
	" pepper jelly
	" pepper sauce
	" pepper taste
	" pepper vodka
	" pepperoncini
	" pequin chile
	" percent milk
	" petrale sole
	" phyllo dough
	" phyllo pastry ",
	" phyllo sheet
	" picante salsa ",
	" picante sauce ",
	" pickle brine
	" pickle juice
	" pickle relish ",
	" pickled beet
	" pickled onion ",
	" pickled pear
	" pickling lime ",
	" pickling salt ",
	" pineapple rum ",
	" pippin apple
	" pistachio nut ",
	" pistachio oil ",
	" pitum pocket
	" plastic wrap
	" plum tomatoes ",
	" plymouth gin
	" poblano chile ",
	" poblano chili ",
	" poire william ",
	" pork sausage
	" pork sparerib ",
	" porridge oat
	" potato bread
	" potato flour
	" potato salad
	" potato starch ",
	" premium vodka ",
	" pretzel salt
	" prickly pear
	" puffed millet ",
	" pullman bread ",
	" pumpernickel
	" pumpkin flesh ",
	" pumpkin puree ",
	" pumpkin seed
	" pureed banana ",
	" pureed tomato ",
	" purple basil
	" purple onion
	" purple potato ",
	" quark cheese
	" queso blanco
	" queso fresco
	" quick oatmeal ",
	" quinoa flour
	" radish green
	" radish sprout ",
	" rainbow chard ",
	" rainbow trout ",
	" raisin bread
	" raman noodle
	" rapeseed oil
	" raspberry jam ",
	" real vanilla
	" refried bean
	" rendered duck ",
	" rhubarb syrup ",
	" rice cracker
	" rice vinegar
	" riesling wine ",
	" risotto rice
	" ritz cracker
	" roast chicken ",
	" roast turkey
	" roasting pan
	" roll wrapper
	" rolled barley ",
	" romaine heart ",
	" romano cheese ",
	" romesco sauce ",
	" rose essence
	" rotel tomato
	" russet potato ",
	" safflower oil ",
	" sage sausage
	" salmon filet
	" salmon steak
	" salsa fresca
	" salsa picante ",
	" sambal oelek
	" sanding sugar ",
	" sandwich bun
	" sandwich roll ",
	" sausage link
	" sausage meat
	" savoy cabbage ",
	" scalded milk
	" scotch bonnet ",
	" scotch whisky ",
	" scrambled egg ",
	" season blend
	" seltzer water ",
	" serrano chile ",
	" serrano chili ",
	" sesame seeds
	" sesame tahini ",
	" shaoxing wine ",
	" sharp cheddar ",
	" shelling bean ",
	" sichuan pepper ",
	" sifted flour
	" single cream
	" sirloin beef
	" sirloin steak ",
	" sirloin strip ",
	" skimmed milk
	" smoked almond ",
	" smoked bacon
	" smoked gouda
	" smoked salmon ",
	" smoked trout
	" smoked turkey ",
	" snack cracker ",
	" soda cracker
	" soft caramel
	" soppressatum
	" sorghum flour ",
	" sorghum syrup ",
	" soy margarine ",
	" spanish olive ",
	" spanish onion ",
	" spanish rice
	" spicy mustard ",
	" spicy sausage ",
	" spiny lobster ",
	" splash water
	" spring garlic ",
	" spring green
	" spring onion
	" spring water
	" squeeze lemon ",
	" squeeze lime
	" steamed rice
	" stewed tomato ",
	" stewing beef
	" stewing lamb
	" strawberries
	" streaky bacon ",
	" string cheese ",
	" striped bass
	" stuffing mix
	" sturdy green
	" sugar cookie
	" sugar needed
	" sugar pumpkin ",
	" sunflower oil ",
	" swiss cheese
	" tabasco pepper ",
	" tabasco sauce ",
	" tahini sauce
	" tamari sauce
	" tamarind pulp ",
	" tapioca flour ",
	" tapioca pearl ",
	" tartar sauce
	" tempeh bacon
	" thick yogurt
	" thyme branch
	" thyme leaves
	" thyme oregano ",
	" thyme parsley ",
	" tiger shrimp
	" tofu crumble
	" tomato basil
	" tomato juice
	" tomato paste
	" tomato pesto
	" tomato puree
	" tomato relish ",
	" tomato salsa
	" tomato sauce
	" tortilla chip ",
	" truffle salt
	" turkey bacon
	" turkey broth
	" turkey gravy
	" turkey thigh
	" turnip green
	" vanilla bean
	" vanilla sauce ",
	" vanilla sugar ",
	" vanilla syrup ",
	" vanilla vodka ",
	" vanilla wafer ",
	" veal scallop
	" vegan butter
	" vegan cheese
	" vegan chicken ",
	" vegetable oil ",
	" veggie burger ",
	" venison steak ",
	" vialone nano
	" vinegar sauce ",
	" virginium ham ",
	" voatsiperifery ",
	" vodka chilled ",
	" wafer cookie
	" walnut butter ",
	" walnut flour
	" walnut pecan
	" water chestnut ",
	" water cracker ",
	" wheat couscou ",
	" wheat cracker ",
	" wheat gluten
	" whipped cream ",
	" white cabbage ",
	" white cheddar ",
	" white cheese
	" white creamer ",
	" white hominy
	" white mustards ",
	" white pepper
	" white potato
	" white quinoa
	" white radish
	" white raisin
	" white tequila ",
	" white truffle ",
	" white turnip
	" white vinegar ",
	" wild arugula
	" wild mushroom ",
	" wine vinegar
	" winter squash ",
	" wondra flour
	" wooden skewer ",
	" yellow apple
	" yellow chile
	" yellow hominy ",
	" yellow lentil ",
	" yellow onion
	" yellow pepper ",
	" yellow potato ",
	" yellow squash ",
	" yellow tomato ",
	" yogurt sauce
	" young carrot
	" achiote seed ",
	" acorn squash ",
	" acorn squashs ",
	" adobo sauce
	" adzuki bean
	" agave nectar ",
	" agave syrup
	" aged cheddar ",
	" ajwain seed
	" alaskan cod
	" aleppo chili ",
	" almond flour ",
	" almond meal
	" almond milk
	" almond skin
	" almond syrup ",
	" amontillado
	" ancho chile
	" ancho chili
	" annatto oil
	" annatto seed ",
	" apple brandy ",
	" apple butter ",
	" apple cider
	" apple jelly
	" apple juice
	" apple puree
	" apple sauce
	" apple syrup
	" apricot jam
	" arbol chile
	" arborio rice ",
	" arctic char
	" arepa flour
	" assorted nut ",
	" averna amaro ",
	" avocado dip
	" avocado oil
	" baby arugula ",
	" baby carrot
	" baby fennel
	" baby lettuce ",
	" baby potato
	" baby rocket
	" baby shrimp
	" baby spinach ",
	" baby squash
	" baby tomato
	" baby turnip
	" bacardi rum
	" bacon grease ",
	" bacon strip
	" bag spinach
	" baked potato ",
	" bakery bread ",
	" baking apple ",
	" baking soda
	" baking spray ",
	" banana bread ",
	" banana chip
	" banana puree ",
	" banana squash ",
	" barbecue rub ",
	" barley flour ",
	" barley malt
	" basil chive
	" basil leaves ",
	" basil pesto
	" basmati rice ",
	" bay scallop
	" bean sprout
	" bean thread
	" beef brisket ",
	" beef consomm ",
	" beef marrow
	" beef sirloin ",
	" beef tongue
	" bell pepper
	" benedictine
	" besan flour
	" beurre blanc ",
	" bibb lettuce ",
	" bing cherry
	" bird pepper
	" biscuit mix
	" black beans
	" black mustard ",
	" black olives ",
	" black pepper ",
	" blackcurrants ",
	" blended oil
	" bleu cheese
	" blood orange ",
	" blood oranges ",
	" blue cheese
	" blue curacao ",
	" blueberries
	" bone removed ",
	" borlotti bean ",
	" bourbon rye
	" boysenberry
	" bread crumb
	" bread crumbs ",
	" bread flour
	" breadcrumbs
	" brie cheese
	" brioche bun
	" brioche roll ",
	" broccoflowers ",
	" broken pecan ",
	" brook trout
	" brown butter ",
	" brown lentil ",
	" brown mustard ",
	" brown onion
	" brown raisin ",
	" brown sugar
	" bruschettum
	" bulb fennel
	" bulb garlic
	" bulb shallot ",
	" bulgar wheat ",
	" bulgur wheat ",
	" bulk sausage ",
	" butter bean
	" butter pecan ",
	" butter sauce ",
	" butterscotch ",
	" californium
	" calrose rice ",
	" can chickpea ",
	" canary melons ",
	" candied peel ",
	" canning salt ",
	" caper berry
	" caper brine
	" caper juice
	" caraway seed ",
	" cardamom pod ",
	" carne asada
	" carrot curl
	" carrot juice ",
	" cashew cream ",
	" cashew half
	" cashew milk
	" caster sugar ",
	" cauliflower
	" celery heart ",
	" celery powder ",
	" celery salt
	" celery seed
	" chaat masala ",
	" chanterelle
	" chat masala
	" cheese curd
	" cheese sauce ",
	" cheese soup
	" cheese whiz
	" cheesecloth
	" cherry cola
	" cherry juice ",
	" cherry syrup ",
	" cherry vodka ",
	" chex cereal
	" chicken bone ",
	" chicken fat
	" chicken meat ",
	" chicken skin ",
	" chicken soup ",
	" chicken wing ",
	" chile pepper ",
	" chile salsa
	" chile sauce
	" chili pepper ",
	" chili peppers ",
	" chili powder ",
	" chili sauce
	" chilled lard ",
	" chimichurri
	" chinese wine ",
	" chipotle rub ",
	" chuck roast
	" chuck steak
	" chunky salsa ",
	" citric acid
	" citru fruit
	" citru juice
	" citru vodka
	" clam chowder ",
	" clear honey
	" clover honey ",
	" cocoa butter ",
	" coconut chip ",
	" coconut meat ",
	" coconut milk ",
	" coconut oil
	" coconut rum
	" coffee bean
	" coffee cake
	" coho salmon
	" colby cheese ",
	" coleslaw mix ",
	" collard green ",
	" comice pear
	" comt cheese
	" comte cheese ",
	" cookie crumb ",
	" cooking oil
	" corn kernel
	" corn kernels ",
	" corn niblet
	" corn relish
	" corn shucked ",
	" corn starch
	" corned beef
	" cornish hen
	" corona beer
	" country ham
	" country loaf ",
	" cranberries
	" cream cheese ",
	" cream sauce
	" cream sherry ",
	" creamed corn ",
	" creole spice ",
	" crisp apple
	" crisp bacon
	" crme frache
	" crusty bread ",
	" crusty roll
	" curing salt
	" curly endive ",
	" curry powder ",
	" custard mix
	" cut chicken
	" deli turkey
	" demerara rum ",
	" dill parsley ",
	" dill pickle
	" dinner roll
	" double cream ",
	" dry mustard
	" dry red wine ",
	" duck confit
	" edam cheese
	" edamame bean ",
	" egg noodles
	" egg whisked
	" english pea
	" fennel bulb
	" fennel frond ",
	" fennel seed
	" feta cheese
	" fetum cheese ",
	" filet mignon ",
	" filo pastry
	" fino sherry
	" flank steak
	" flat anchovy ",
	" flat parsley ",
	" flaxseed oil ",
	" floral honey ",
	" frankfurter
	" french bean
	" french bread ",
	" french roll
	" french toast ",
	" fresno chile ",
	" fried onion
	" fried potato ",
	" fruit juice
	" fruit salad
	" fruit salsa
	" gaetum olive ",
	" garam masala ",
	" garlic broth ",
	" garlic bulb
	" garlic chive ",
	" garlic chives ",
	" garlic hummu ",
	" garlic juice ",
	" garlic salt
	" garlic scape ",
	" gefilte fish ",
	" gem lettuce
	" giblet broth ",
	" gigante bean ",
	" ginger beer
	" ginger juice ",
	" ginger snap
	" ginger syrup ",
	" gingerbread
	" globe tomato ",
	" gluten flour ",
	" gluten free
	" goat cheese
	" gold tequila ",
	" golden beet
	" golden syrup ",
	" gouda cheese ",
	" grana padano ",
	" granola bar
	" grape jelly
	" grape tomato ",
	" greek fetum
	" greek olive
	" greek yogurt ",
	" green apple
	" green banana ",
	" green beans
	" green cherry ",
	" green chile
	" green chili
	" green garlic ",
	" green grape
	" green lentil ",
	" green mango
	" green olive
	" green onion
	" green onions ",
	" green papaya ",
	" green pepper ",
	" green peppers ",
	" green salsa
	" green tomato ",
	" guava juice
	" guava nectar ",
	" haa avocado
	" halved lemon ",
	" hanger steak ",
	" hardy green
	" haricot vert ",
	" hazelnut oil ",
	" hearty bread ",
	" hearty green ",
	" heavy cream
	" herb bouquet ",
	" herb butter
	" herb mixture ",
	" hickory chip ",
	" hoisin sauce ",
	" hollandaise
	" hominy grit
	" honey syrup
	" horseradish
	" huckleberry
	" icing sugar
	" inch celery
	" inch ginger
	" instant grit ",
	" instant milk ",
	" iodized salt ",
	" irish cream
	" italian herb ",
	" italian plum ",
	" italian tuna ",
	" jack cheese
	" jamaica rum
	" jamaican rum ",
	" jarred salsa ",
	" jasmine rice ",
	" jigger vodka ",
	" juice orange ",
	" jumbo shrimp ",
	" juniper berry ",
	" kaffir lime
	" kaiser roll
	" kernel corn
	" ketjap mani
	" kidney bean
	" kidney beans ",
	" kimchi juice ",
	" king salmon
	" kirschwasser ",
	" kosher salt
	" langoustine
	" laughing cow ",
	" lavender bud ",
	" leafy green
	" lemon aioli
	" lemon basil
	" lemon grass
	" lemon halved ",
	" lemon juice
	" lemon myrtles ",
	" lemon pepper ",
	" lemon thyme
	" lemon verbena ",
	" lemon vodka
	" lemon wheel
	" lemon yogurt ",
	" lillet blanc ",
	" lime butter
	" lime halved
	" lingonberry
	" link sausage ",
	" litsea cubeba ",
	" little water ",
	" live lobster ",
	" live mussel
	" loaf brioche ",
	" loaf challah ",
	" lobster meat ",
	" lobster tail ",
	" london broil ",
	" lowfat milk
	" lychee juice ",
	" madeira wine ",
	" madra curry
	" maldon salt
	" malt vinegar ",
	" mango puree
	" mango salsa
	" mango sorbet ",
	" mango-gingers ",
	" manila clam
	" maple sugar
	" maple syrup
	" mara pepper
	" marrow bone
	" marsala wine ",
	" marshmallow
	" marshmallows ",
	" masa harina
	" matzo farfel ",
	" matzoh meal
	" meat filling ",
	" meatloaf mix ",
	" medjool date ",
	" metal skewer ",
	" mexican beer ",
	" mexican rice ",
	" meyer lemon
	" milk yogurt
	" millet flour ",
	" mint garnish ",
	" mint leaves
	" mint parsley ",
	" minute rice
	" miracle whip ",
	" mixed berry
	" mixed citru
	" mixed fruit
	" mixed green
	" mixed olive
	" mixed spice
	" mixed tomato ",
	" muscat wine
	" mushroom cap ",
	" mustard green ",
	" mustard oil
	" mustard seed ",
	" napa cabbage ",
	" navel orange ",
	" neutral oil
	" nigella seed ",
	" nilla wafer
	" noilly prat
	" nonfat milk
	" nori seaweed ",
	" nut biscuit
	" onion powder ",
	" orange juice ",
	" orange peel
	" orange sauce ",
	" orange soda
	" orange wheel ",
	" orecchiette
	" oreo cookie
	" orgeat syrup ",
	" oyster sauce ",
	" oz parmesan
	" oz prosecco
	" oz spaghetti ",
	" pacific cod
	" paella rice
	" pan dripping ",
	" pancake mix
	" panch phoron ",
	" papaya seed
	" pappardelle
	" parsley dill ",
	" parsley mint ",
	" passionfruits ",
	" pastry brush ",
	" pastry cream ",
	" pastry dough ",
	" pastry flour ",
	" pea tendril
	" peach brandy ",
	" peach nectar ",
	" peach puree
	" peach syrup
	" peach yogurt ",
	" peanut sauce ",
	" pear brandy
	" pear liqueur ",
	" pear nectar
	" pear tomato
	" pearl barley ",
	" pearl onion
	" pearl sugar
	" pecan walnut ",
	" peking duck
	" penne rigate ",
	" peperoncini
	" peperoncino
	" pepper jelly ",
	" pepper sauce ",
	" pepper taste ",
	" pepper vodka ",
	" pepperoncini ",
	" pequin chile ",
	" percent milk ",
	" perciatelli
	" pesto sauce
	" petrale sole ",
	" phyllo dough ",
	" phyllo sheet ",
	" pickle brine ",
	" pickle chip
	" pickle juice ",
	" pickled beet ",
	" pickled pear ",
	" pig trotter
	" pignoli nut
	" pippin apple ",
	" pitum bread
	" pitum pocket ",
	" pizza crust
	" pizza dough
	" pizza sauce
	" plastic wrap ",
	" plum tomato
	" plymouth gin ",
	" poached egg
	" pomegranate
	" ponzu sauce
	" popped corn
	" pork sausage ",
	" porridge oat ",
	" potato bread ",
	" potato chip
	" potato flour ",
	" potato roll
	" potato salad ",
	" pretzel salt ",
	" prickly pear ",
	" prune juice
	" prune puree
	" pudding mix
	" puff pastry
	" puffed rice
	" pumpernickel ",
	" pumpkin pie
	" pumpkin seed ",
	" pumpkinseed
	" purple basil ",
	" purple onion ",
	" purple plum
	" quark cheese ",
	" quatre pice
	" queso blanco ",
	" queso fresco ",
	" quick bread
	" quinoa flour ",
	" radish green ",
	" raisin bread ",
	" raman noodle ",
	" rapeseed oil ",
	" raspberries
	" real butter
	" real vanilla ",
	" red peppers
	" refried bean ",
	" rice cereal
	" rice cracker ",
	" rice krispy
	" rice noodle
	" rice vinegar ",
	" risotto rice ",
	" rittenhouse
	" ritz cracker ",
	" roast turkey ",
	" roasting pan ",
	" roll wrapper ",
	" roma tomato
	" romano bean
	" rose essence ",
	" rotel tomato ",
	" royal icing
	" rubbed sage
	" rum liqueur
	" rye whiskey
	" sage sausage ",
	" salad green
	" salmon filet ",
	" salmon skin
	" salmon steak ",
	" salsa fresca ",
	" salsa verde
	" salt pepper
	" sambal oelek ",
	" san marzano
	" sandwich bun ",
	" sarsaparillas ",
	" satay sauce
	" sausage link ",
	" sausage meat ",
	" scalded milk ",
	" season blend ",
	" season salt
	" seckel pear
	" serrano ham
	" sesame seed
	" sesame seeds ",
	" shallot oil
	" sheet matzo
	" sherry wine
	" sifted flour ",
	" silken tofu
	" single cream ",
	" sirloin beef ",
	" sirloin tip
	" skimmed milk ",
	" skirt steak
	" smith apple
	" smoked bacon ",
	" smoked fish
	" smoked gouda ",
	" smoked salt
	" smoked tofu
	" smoked trout ",
	" smoky bacon
	" soba noodle
	" soda cracker ",
	" soft butter
	" soft caramel ",
	" soft cheese
	" soppressatum ",
	" sour cherry
	" sour pickle
	" soy protein
	" soybean oil
	" spaghettini
	" spanish rice ",
	" spelt berry
	" spelt flour
	" spicy green
	" splash water ",
	" sponge cake
	" spring green ",
	" spring onion ",
	" spring water ",
	" squash seed
	" squeeze lime ",
	" stale bread
	" steak sauce
	" steamed rice ",
	" stewing beef ",
	" stewing lamb ",
	" sticky rice
	" strawberries ",
	" string bean
	" strip bacon
	" strip steak
	" striped bass ",
	" stuffing mix ",
	" sturdy green ",
	" sugar cookie ",
	" sugar needed ",
	" sugar syrup
	" sweet onion
	" sweet potatos ",
	" swiss chard
	" swiss cheese ",
	" table cream
	" tagliatelle
	" tahini sauce ",
	" tamari sauce ",
	" tart cherry
	" tartar sauce ",
	" tempeh bacon ",
	" tender herb
	" tepid water
	" thai pepper
	" thick bacon
	" thick yogurt ",
	" thyme branch ",
	" thyme honey
	" thyme leaves ",
	" tiger prawn
	" tiger shrimp ",
	" tiny potato
	" tium marium
	" toast point
	" tofu crumble ",
	" tomato basil ",
	" tomato juice ",
	" tomato paste ",
	" tomato pesto ",
	" tomato puree ",
	" tomato salsa ",
	" tomato sauce ",
	" tomato soup
	" tonic water
	" tortiglioni
	" truffle oil
	" truffle salt ",
	" turkey bacon ",
	" turkey bone
	" turkey broth ",
	" turkey gravy ",
	" turkey meat
	" turkey thigh ",
	" turkey wing
	" turnip green ",
	" udon noodle
	" vanilla bean ",
	" vanilla pod
	" veal scallop ",
	" vegan butter ",
	" vegan cheese ",
	" vegan sugar
	" vialone nano ",
	" vinaigrette
	" vsop cognac
	" wafer cookie ",
	" walnut flour ",
	" walnut half
	" walnut pecan ",
	" waxy potato
	" wheat bagel
	" wheat berry
	" wheat bread
	" wheat flour
	" wheat gluten ",
	" wheat penne
	" wheat pitum
	" wheat toast
	" white cheese ",
	" white flour
	" white grape
	" white hominy ",
	" white mustard ",
	" white onion
	" white peach
	" white pepper ",
	" white potato ",
	" white quinoa ",
	" white radish ",
	" white radishs ",
	" white raisin ",
	" white sauce
	" white sugar
	" white turnip ",
	" white wheat
	" wild arugula ",
	" wild salmon
	" wine vinegar ",
	" wondra flour ",
	" wonton skin
	" xanthan gum
	" yellow apple ",
	" yellow beet
	" yellow chile ",
	" yellow corn
	" yellow finn
	" yellow grit
	" yellow miso
	" yellow onion ",
	" yellow rice
	" yogurt sauce ",
	" young carrot ",
	" acorn squash ",
	" adobo sauce ",
	" adzuki bean ",
	" agave syrup ",
	" ajwain seed ",
	" alaskan cod ",
	" almond meal ",
	" almond milk ",
	" almond oil
	" almond skin ",
	" amber beer
	" amontillado ",
	" ancho chile ",
	" ancho chili ",
	" anise seed
	" anjou pear
	" annatto oil ",
	" apple cake
	" apple cider ",
	" apple jelly ",
	" apple juice ",
	" apple pear
	" apple puree ",
	" apple sauce ",
	" apple syrup ",
	" applesauce
	" apricot jam ",
	" arbol chile ",
	" arctic char ",
	" arepa flour ",
	" asafoetida
	" asian pear
	" avocado dip ",
	" avocado oil ",
	" baby caper
	" baby carrot ",
	" baby fennel ",
	" baby green
	" baby onion
	" baby potato ",
	" baby rocket ",
	" baby shrimp ",
	" baby squash ",
	" baby tomato ",
	" baby turnip ",
	" bacardi rum ",
	" bacon strip ",
	" bag spinach ",
	" bagel chip
	" baked bean
	" baked tofu
	" baking mix
	" baking soda ",
	" banana chip ",
	" banana rum
	" barley malt ",
	" basil chive ",
	" basil mint
	" basil pesto ",
	" bay leaves
	" bay scallop ",
	" bay shrimp
	" bean broth
	" bean puree
	" bean salad
	" bean sprout ",
	" bean sprouts ",
	" bean thread ",
	" beaten egg
	" beef broth
	" beef cheek
	" beef chuck
	" beef liver
	" beef marrow ",
	" beef mince
	" beef roast
	" beef steak
	" beef stock
	" beef tongue ",
	" beet green
	" bell pepper ",
	" bell peppers ",
	" benedictine ",
	" benne seed
	" berry wine
	" besan flour ",
	" big carrot
	" bing cherry ",
	" bird chile
	" bird pepper ",
	" biscuit mix ",
	" black beans ",
	" blackberry
	" blackcurrant ",
	" blended oil ",
	" bleu cheese ",
	" blood orange ",
	" blue cheese ",
	" blueberries ",
	" bocconcini
	" bomba rice
	" bourbon rye ",
	" boysenberry ",
	" boysenberrys ",
	" brazil nut
	" bread crumb ",
	" bread flour ",
	" bread roll
	" breadcrumb
	" breadcrumbs ",
	" breadfruit
	" brie cheese ",
	" brioche bun ",
	" broad bean
	" broad beanss ",
	" broccoflower ",
	" broccolini
	" brook trout ",
	" brown onion ",
	" brown rice
	" brown sugar ",
	" bruschettum ",
	" bulb fennel ",
	" bulb garlic ",
	" burger bun
	" butter bean ",
	" butter oil
	" buttermilk
	" cake flour
	" cake yeast
	" calf liver
	" californium ",
	" campanelle
	" canary melon ",
	" candy cane
	" cane sugar
	" cane syrup
	" cannellini
	" cannelloni
	" canola oil
	" cantaloupe
	" caper berry ",
	" caper brine ",
	" caper juice ",
	" caperberry
	" cappuccino
	" carne asada ",
	" carob chip
	" carrot curl ",
	" cashew half ",
	" cashew milk ",
	" cashew nut
	" cauliflower ",
	" cauliflowers ",
	" celery rib
	" celery salt ",
	" celery seed ",
	" celery seeds ",
	" channa dal
	" chanterelle ",
	" chardonnay
	" chat masala ",
	" cheese curd ",
	" cheese soup ",
	" cheese whiz ",
	" cheesecake
	" cheesecloth ",
	" cherry cola ",
	" cherry jam
	" chex cereal ",
	" chicken fat ",
	" chile salsa ",
	" chile sauce ",
	" chili bean
	" chili pepper ",
	" chili sauce ",
	" chimichurri ",
	" chium seed
	" chive dill
	" chuck roast ",
	" chuck steak ",
	" citric acid ",
	" citru fruit ",
	" citru juice ",
	" citru vodka ",
	" clam broth
	" clam juice
	" clear honey ",
	" clementine
	" co lettuce
	" coconut oil ",
	" coconut rum ",
	" coffee bean ",
	" coffee cake ",
	" coho salmon ",
	" comice pear ",
	" comt cheese ",
	" cooking oil ",
	" cool water
	" corn bread
	" corn flour
	" corn kernel ",
	" corn niblet ",
	" corn relish ",
	" corn starch ",
	" corn syrup
	" corned beef ",
	" cornish hen ",
	" cornstarch
	" corona beer ",
	" cottonseed
	" country ham ",
	" cranberries ",
	" cream sauce ",
	" crisp apple ",
	" crisp bacon ",
	" crme frache ",
	" crusty roll ",
	" cumin seed
	" cured meat
	" curing salt ",
	" curly kale
	" custard mix ",
	" cut chicken ",
	" cuttlefish
	" date sugar
	" date syrup
	" decoration
	" deli turkey ",
	" demi glace
	" dill chive
	" dill frond
	" dill pickle ",
	" dilly bean
	" dinner roll ",
	" dragonfruits ",
	" dry mustard ",
	" dry sherry
	" duck confit ",
	" edam cheese ",
	" egg noodle
	" egg noodles ",
	" egg tomato
	" egg whisked ",
	" elderberry
	" emmentaler
	" english pea ",
	" erythritol
	" farfalline
	" fennel bulb ",
	" fennel seed ",
	" feta cheese ",
	" fettuccine
	" filo pastry ",
	" fino sherry ",
	" fish broth
	" fish sauce
	" fish steak
	" flaky salt
	" flank steak ",
	" flat bread
	" frangelico
	" frankfurter ",
	" french bean ",
	" french fry
	" french roll ",
	" fried onion ",
	" fried rice
	" fried tofu
	" fruit juice ",
	" fruit salad ",
	" fruit salsa ",
	" frying oil
	" gala apple
	" garden pea
	" garlic bulb ",
	" garlic oil
	" garlic salt ",
	" gem lettuce ",
	" ginger ale
	" ginger beer ",
	" ginger snap ",
	" gingerbread ",
	" gingerroot
	" gingersnap
	" gluten free ",
	" goat cheese ",
	" goji berry
	" golden beet ",
	" golden rum
	" good bread
	" good honey
	" gooseberry
	" gorgonzola
	" granola bar ",
	" grape jelly ",
	" grapefruit
	" greek fetum ",
	" greek olive ",
	" green apple ",
	" green bean
	" green beans ",
	" green chile ",
	" green chili ",
	" green grape ",
	" green mango ",
	" green olive ",
	" green onion ",
	" green pear
	" green pepper ",
	" green salsa ",
	" guava juice ",
	" gumbo file
	" haa avocado ",
	" hardy green ",
	" hash brown
	" heavy cream ",
	" herb butter ",
	" hoagie bun
	" hollandaise ",
	" holy basil
	" hominy grit ",
	" honey syrup ",
	" horseradish ",
	" horseradishs ",
	" huckleberry ",
	" huckleberrys ",
	" iced water
	" icing sugar ",
	" inch celery ",
	" inch ginger ",
	" irish cream ",
	" jack cheese ",
	" jamaica rum ",
	" jelly bean
	" kaffir lime ",
	" kaiser roll ",
	" kecap mani
	" kernel corn ",
	" ketjap mani ",
	" kidney bean ",
	" kidney beans ",
	" king salmon ",
	" kiwi fruit
	" kosher salt ",
	" lady apple
	" ladyfinger
	" lager beer
	" lamb mince
	" lamb steak
	" langoustine ",
	" leafy green ",
	" leek green
	" leek white
	" lemon aioli ",
	" lemon balm
	" lemon basil ",
	" lemon curd
	" lemon grass ",
	" lemon grasss ",
	" lemon half
	" lemon juice ",
	" lemon lime
	" lemon myrtle ",
	" lemon peel
	" lemon rind
	" lemon sole
	" lemon thyme ",
	" lemon vodka ",
	" lemon wheel ",
	" lemongrass
	" lime butter ",
	" lime halved ",
	" lime juice
	" lime wheel
	" limoncello
	" lingonberry ",
	" live mussel ",
	" liver pate
	" liverwurst
	" loaf bread
	" long peppers ",
	" lowfat milk ",
	" madra curry ",
	" maldon salt ",
	" malibu rum
	" malt syrup
	" mango pulp
	" mango puree ",
	" mango salsa ",
	" mango-ginger ",
	" manila clam ",
	" maple sugar ",
	" maple syrup ",
	" mara pepper ",
	" margaritum
	" marrow bone ",
	" marshmallow ",
	" masa harina ",
	" mascarpone
	" matzo meal
	" matzoh meal ",
	" mayonnaise
	" meat sauce
	" meyer lemon ",
	" microgreen
	" milk cream
	" milk crumb
	" milk yogurt ",
	" minestrone
	" mint leaves ",
	" mint pesto
	" mint sauce
	" mint syrup
	" minute rice ",
	" mixed berry ",
	" mixed citru ",
	" mixed fruit ",
	" mixed green ",
	" mixed herb
	" mixed olive ",
	" mixed seed
	" mixed spice ",
	" mole sauce
	" mortadella
	" mozzarella
	" muscat wine ",
	" mustard oil ",
	" naan bread
	" neutral oil ",
	" new potato
	" nilla wafer ",
	" noilly prat ",
	" nonfat milk ",
	" nori sheet
	" nut biscuit ",
	" nut butter
	" oil butter
	" onion salt
	" opal basil
	" orange oil
	" orange peel ",
	" orange soda ",
	" orecchiette ",
	" oreo cookie ",
	" oz parmesan ",
	" oz prosecco ",
	" pacific cod ",
	" paella rice ",
	" pale lager
	" palm sugar
	" pancake mix ",
	" papaya seed ",
	" pappardelle ",
	" parmigiano
	" passionfruit ",
	" pat butter
	" pea tendril ",
	" peach puree ",
	" peach syrup ",
	" peanut oil
	" pear brandy ",
	" pear juice
	" pear nectar ",
	" pear tomato ",
	" pearl onion ",
	" pearl sugar ",
	" pecan meal
	" pekin duck
	" peking duck ",
	" peperoncini ",
	" peperoncino ",
	" peppercorn
	" peppermint
	" perciatelli ",
	" pesto sauce ",
	" petite pea
	" pickle chip ",
	" pie pastry
	" pig trotter ",
	" pigeon pea
	" pignoli nut ",
	" piloncillo
	" pinot noir
	" pinto bean
	" pitum bread ",
	" pizza crust ",
	" pizza dough ",
	" pizza sauce ",
	" plum sauce
	" plum tomato ",
	" poached egg ",
	" pomegranate ",
	" pomegranates ",
	" ponzu sauce ",
	" popped corn ",
	" poppy seed
	" pork belly
	" pork liver
	" pork roast
	" pork steak
	" portobello
	" potato bun
	" potato chip ",
	" potato roll ",
	" prosciutto
	" prune juice ",
	" prune plum
	" prune puree ",
	" pudding mix ",
	" puff pastry ",
	" puffed rice ",
	" pumpkin pie ",
	" pumpkinseed ",
	" purple plum ",
	" quaker oat
	" quatre pice ",
	" quesadilla
	" quick bread ",
	" ramp green
	" raspberries ",
	" real butter ",
	" red pepper
	" red peppers ",
	" rib celery
	" rice cereal ",
	" rice flour
	" rice krispy ",
	" rice noodle ",
	" rice paper
	" rice pilaf
	" rittenhouse ",
	" roast beef
	" roast pork
	" rolled oat
	" roma tomato ",
	" romano bean ",
	" rome apple
	" rose petal
	" rose water
	" royal icing ",
	" rubbed sage ",
	" rum liqueur ",
	" rump roast
	" rump steak
	" runner beans ",
	" rye whiskey ",
	" rye whisky
	" sage thyme
	" salad green ",
	" salal berrys ",
	" salmon roe
	" salmon skin ",
	" salsa verde ",
	" salt pepper ",
	" salt water
	" san marzano ",
	" sarsaparilla ",
	" satay sauce ",
	" sauerkraut
	" season salt ",
	" seckel pear ",
	" serrano ham ",
	" sesame oil
	" sesame seed ",
	" shallot oil ",
	" sheet matzo ",
	" sherry wine ",
	" shiro miso
	" shortbread
	" shortening
	" silken tofu ",
	" sirloin tip ",
	" skate wing
	" skirt steak ",
	" slider bun
	" smith apple ",
	" smoked fish ",
	" smoked ham
	" smoked salt ",
	" smoked tofu ",
	" smoky bacon ",
	" snack cake
	" soba noodle ",
	" soda bread
	" soda water
	" soft butter ",
	" soft cheese ",
	" soft drink
	" soft fruit
	" solid tuna
	" sour apple
	" sour cherry ",
	" sour cream
	" sour pickle ",
	" soy protein ",
	" soy yogurt
	" soya sauce
	" soybean oil ",
	" spaghettini ",
	" spelt berry ",
	" spelt flour ",
	" spiced rum
	" spicy green ",
	" sponge cake ",
	" squash seed ",
	" st germain
	" stale bread ",
	" star anise
	" star fruit
	" steak sauce ",
	" sticky rice ",
	" stout beer
	" strawberry
	" string bean ",
	" strip bacon ",
	" strip steak ",
	" stroganoff
	" sugar cane
	" sugar cone
	" sugar syrup ",
	" sushi rice
	" sweet onion ",
	" sweet potato ",
	" sweetbread
	" swiss chard ",
	" table cream ",
	" table salt
	" tagliatelle ",
	" tart apple
	" tart cherry ",
	" tart crust
	" tawny port
	" teff flour
	" tender herb ",
	" tepid water ",
	" thai basil
	" thai chile
	" thai chili
	" thai pepper ",
	" thick bacon ",
	" thyme honey ",
	" tiger prawn ",
	" tiny caper
	" tiny potato ",
	" tium marium ",
	" toast point ",
	" toffee bit
	" tomato jam
	" tomato oil
	" tomato soup ",
	" tonic water ",
	" torn basil
	" torn green
	" tortellini
	" tortiglioni ",
	" triple sec
	" truffle oil ",
	" tuna salad
	" tuna steak
	" tuna water
	" turkey bone ",
	" turkey fat
	" turkey ham
	" turkey meat ",
	" turkey wing ",
	" udon noodle ",
	" vanilla pod ",
	" veal roast
	" vegan sugar ",
	" vermicelli
	" vinaigrette ",
	" vsop cognac ",
	" walnut half ",
	" walnut oil
	" watercress
	" watermelon
	" waxy potato ",
	" weisswurst
	" wheat bagel ",
	" wheat beer
	" wheat berry ",
	" wheat bran
	" wheat bread ",
	" wheat flour ",
	" wheat germ
	" wheat penne ",
	" wheat pita
	" wheat pitum ",
	" wheat toast ",
	" white bean
	" white beer
	" white corn
	" white fish
	" white flour ",
	" white grape ",
	" white grit
	" white karo
	" white miso
	" white onion ",
	" white peach ",
	" white port
	" white radish ",
	" white rice
	" white sauce ",
	" white sugar ",
	" white tuna
	" white wheat ",
	" white wine
	" wild salmon ",
	" wonton skin ",
	" xanthan gum ",
	" yellow beet ",
	" yellow corn ",
	" yellow finn ",
	" yellow grit ",
	" yellow miso ",
	" yellow rice ",
	" yukon gold
	" yuzu juice
	" ziti penne
	" agar agar
	" almond oil ",
	" amber beer ",
	" amber rum
	" andouille
	" anise oil
	" anise seed ",
	" anjou pear ",
	" apple cake ",
	" apple jam
	" apple pear ",
	" apple pie
	" applejack
	" applesauce ",
	" arrowroot
	" artichoke
	" asafetida
	" asafoetida ",
	" asafoetidas ",
	" asian pear ",
	" asparagus
	" aubergine
	" azuki beans ",
	" baby beet
	" baby caper ",
	" baby clam
	" baby corn
	" baby green ",
	" baby kale
	" baby lamb
	" baby leek
	" baby onion ",
	" bacon bit
	" bacon fat
	" bagel chip ",
	" baked bean ",
	" baked ham
	" baked tofu ",
	" baking mix ",
	" banana rum ",
	" basil mint ",
	" basil oil
	" bay leaves ",
	" bay shrimp ",
	" bbq sauce
	" bean broth ",
	" bean puree ",
	" bean salad ",
	" bean sprout ",
	" beaten egg ",
	" beef bone
	" beef broth ",
	" beef cheek ",
	" beef chuck ",
	" beef liver ",
	" beef mince ",
	" beef roast ",
	" beef steak ",
	" beef stew
	" beef stock ",
	" beef suet
	" beefsteak
	" beet green ",
	" bell pepper ",
	" benne seed ",
	" berry wine ",
	" big carrot ",
	" bird chile ",
	" black beans ",
	" blackberry ",
	" blackberrys ",
	" blackfish
	" blue crab
	" blueberry
	" bocconcini ",
	" bolognese
	" bomba rice ",
	" bosc pear
	" boysenberry ",
	" bratwurst
	" brazil nut ",
	" bread roll ",
	" breadcrumb ",
	" breadfruit ",
	" breadfruits ",
	" breakfast
	" broad bean ",
	" broad beans ",
	" broccolini ",
	" brown rice ",
	" buckwheat
	" budweiser
	" burger bun ",
	" butter oil ",
	" buttermilk ",
	" butternut
	" cacao nib
	" cake flour ",
	" cake tofu
	" cake yeast ",
	" calamatum
	" calf liver ",
	" camembert
	" campanelle ",
	" candlenut
	" candy bar
	" candy cane ",
	" cane sugar ",
	" cane syrup ",
	" cannellini ",
	" cannelloni ",
	" canola oil ",
	" cantaloupe ",
	" cantaloupes ",
	" capellini
	" caperberry ",
	" capocollo
	" caponatum
	" cappuccino ",
	" carambola
	" cardamaro
	" caribbean
	" carob chip ",
	" cashew nut ",
	" cauliflower ",
	" cavatappi
	" cavatelli
	" celery rib ",
	" celery seed ",
	" chamomile
	" champagne
	" channa dal ",
	" chardonnay ",
	" cheesecake ",
	" chermoula
	" cherry jam ",
	" chick pea
	" chickpeas
	" chile oil
	" chili bean ",
	" chili oil
	" chium seed ",
	" chive dill ",
	" chive plu
	" chocolate
	" chopstick
	" chow mein
	" ciabattum
	" clam broth ",
	" clam juice ",
	" clementine ",
	" clementines ",
	" cloudberrys ",
	" club soda
	" co lettuce ",
	" coca cola
	" cocoa nib
	" cod steak
	" cointreau
	" condiment
	" cool water ",
	" cool whip
	" coriander
	" corn bread ",
	" corn chip
	" corn flour ",
	" corn grit
	" corn husk
	" corn meal
	" corn salads ",
	" corn syrup ",
	" cornbread
	" cornflake
	" cornflour
	" cornichon
	" cornstarch ",
	" cottonseed ",
	" courgette
	" crab cake
	" crab meat
	" cranberry
	" croissant
	" cumin seed ",
	" cured ham
	" cured meat ",
	" curly kale ",
	" curry leafs ",
	" curry oil
	" cuttlefish ",
	" dandelion
	" date sugar ",
	" date syrup ",
	" decoration ",
	" deli meat
	" demi glace ",
	" dill chive ",
	" dill frond ",
	" dill seed
	" dill weed
	" dilly bean ",
	" dr pepper
	" dragonfruit ",
	" dried limes ",
	" dry sherry ",
	" egg noodle ",
	" egg salad
	" egg tomato ",
	" egg white
	" elderberry ",
	" elderberrys ",
	" emmentaler ",
	" emmenthal
	" enchilada
	" entrecote
	" erythritol ",
	" escabeche
	" everclear
	" farfalline ",
	" fava bean
	" fenugreek
	" fettuccine ",
	" fettucine
	" fiddleheads ",
	" field pea
	" fingerroots ",
	" fish bone
	" fish broth ",
	" fish cake
	" fish sauce ",
	" fish steak ",
	" flaky salt ",
	" flat bread ",
	" flatbread
	" flavoring
	" flax meal
	" flax seed
	" focaccium
	" framboise
	" frangelico ",
	" french fry ",
	" fried egg
	" fried rice ",
	" fried tofu ",
	" frittatum
	" fruit jam
	" frying oil ",
	" gala apple ",
	" garden pea ",
	" garlic oil ",
	" gem squashs ",
	" ginger ale ",
	" gingerroot ",
	" gingersnap ",
	" goat meat
	" goat milk
	" gochugaru
	" gochujang
	" goji berry ",
	" goji berrys ",
	" golden rum ",
	" good bread ",
	" good honey ",
	" goose fat
	" gooseberry ",
	" gooseberrys ",
	" gorgonzola ",
	" grapefruit ",
	" grapefruits ",
	" grapeseed
	" green bean ",
	" green beans ",
	" green pea
	" green pear ",
	" grenadine
	" guacamole
	" guanciale
	" gumbo file ",
	" half half
	" ham steak
	" hamburger
	" hash brown ",
	" hazelnuts
	" heath bar
	" hemp seed
	" herb salt
	" herbsaint
	" hero roll
	" hoagie bun ",
	" hoja santas ",
	" holy basil ",
	" holy basils ",
	" honeycomb
	" horseradish ",
	" hot sauce
	" huckleberry ",
	" iced water ",
	" jackfruit
	" jambalaya
	" jelly bean ",
	" jumbo egg
	" kecap mani ",
	" kidney bean ",
	" kiwi fruit ",
	" kiwi fruits ",
	" kiwifruit
	" kochujang
	" kochukaru
	" lady apple ",
	" ladyfinger ",
	" lager beer ",
	" lamb bone
	" lamb meat
	" lamb mince ",
	" lamb steak ",
	" lamb stew
	" leek green ",
	" leek white ",
	" lemon balm ",
	" lemon balms ",
	" lemon curd ",
	" lemon grass ",
	" lemon half ",
	" lemon lime ",
	" lemon peel ",
	" lemon rind ",
	" lemon sole ",
	" lemongrass ",
	" lima bean
	" lime half
	" lime juice ",
	" lime peel
	" lime wheel ",
	" limoncello ",
	" liver pate ",
	" liverwurst ",
	" loaf bread ",
	" long bean
	" long pepper ",
	" mahi mahi
	" malibu rum ",
	" malt syrup ",
	" mango pulp ",
	" mango rum
	" manicotti
	" margarine
	" margaritum ",
	" marmalade
	" mascarpone ",
	" matzo meal ",
	" mayonnaise ",
	" meat loaf
	" meat sauce ",
	" medallion
	" membrillo
	" microgreen ",
	" milk cream ",
	" milk crumb ",
	" minestrone ",
	" mint pesto ",
	" mint sauce ",
	" mint syrup ",
	" mixed herb ",
	" mixed nut
	" mixed seed ",
	" mole sauce ",
	" mortadella ",
	" mozzarella ",
	" mung bean
	" mushrooms
	" naan bread ",
	" navy bean
	" nectarine
	" new potato ",
	" newspaper
	" nori sheet ",
	" nuoc cham
	" nut butter ",
	" oat flour
	" oil butter ",
	" oil spray
	" olive oil
	" onion salt ",
	" opal basil ",
	" orange oil ",
	" pale lager ",
	" palm sugar ",
	" pan juice
	" pan spray
	" pancettum
	" panettone
	" parchment
	" parma ham
	" parmigiano ",
	" partridge
	" pat butter ",
	" peach jam
	" peanut oil ",
	" pear juice ",
	" pecan meal ",
	" pecan nut
	" pekin duck ",
	" peppercorn ",
	" peppermint ",
	" pepperoni
	" persimmon
	" petite pea ",
	" pie crust
	" pie dough
	" pie pastry ",
	" pie shell
	" pigeon pea ",
	" piloncillo ",
	" pine nuts
	" pineapple
	" pinot noir ",
	" pinto bean ",
	" pinto beans ",
	" piri piri
	" pistachio
	" plum sauce ",
	" plum wine
	" pole bean
	" pomegranate ",
	" poppy seed ",
	" poppy seeds ",
	" pork belly ",
	" pork butt
	" pork lard
	" pork liver ",
	" pork meat
	" pork roast ",
	" pork steak ",
	" port wine
	" portobello ",
	" pot roast
	" potato bun ",
	" prosciutto ",
	" provolone
	" prune plum ",
	" quail egg
	" quaker oat ",
	" quesadilla ",
	" quick oat
	" radiatore
	" radicchio
	" ramp green ",
	" raspberry
	" reblochon
	" red onion
	" red pepper ",
	" redcurrants ",
	" remoulade
	" rib celery ",
	" rib roast
	" rib steak
	" rice cake
	" rice flour ",
	" rice milk
	" rice paper ",
	" rice pilaf ",
	" rice wine
	" roast beef ",
	" roast pork ",
	" rock melons ",
	" rock salt
	" rolled oat ",
	" rome apple ",
	" roquefort
	" rose petal ",
	" rose water ",
	" rose wine
	" rosewater
	" ruby port
	" rump roast ",
	" rump steak ",
	" runner bean ",
	" rye berry
	" rye bread
	" rye flour
	" rye whisky ",
	" sablefish
	" safflower
	" sage thyme ",
	" salad oil
	" salal berry ",
	" salmon roe ",
	" salt pork
	" salt water ",
	" sauerkraut ",
	" savoiardi
	" scallions
	" schnitzel
	" sesame oil ",
	" shellfish
	" shiro miso ",
	" short rib
	" shortbread ",
	" shortcake
	" shortening ",
	" skate wing ",
	" skim milk
	" slider bun ",
	" sloppy jo
	" smoked ham ",
	" snack cake ",
	" snap bean
	" soda bread ",
	" soda water ",
	" soft drink ",
	" soft fruit ",
	" soft herb
	" soft roll
	" soft tofu
	" solid tuna ",
	" sour apple ",
	" sour cream ",
	" sourdough
	" soy cream
	" soy flour
	" soy sauce
	" soy yogurt ",
	" soya milk
	" soya sauce ",
	" spaghetti
	" spare rib
	" spearmint
	" spice mix
	" spice rub
	" spiced rum ",
	" spirulina
	" split pea
	" squid ink
	" st germain ",
	" star anise ",
	" star anises ",
	" star fruit ",
	" star fruits ",
	" stew beef
	" stew meat
	" stout beer ",
	" strawberry ",
	" strawberrys ",
	" stroganoff ",
	" succotash
	" sugar cane ",
	" sugar cone ",
	" sunflower
	" sushi rice ",
	" sweetbread ",
	" sweetcorn
	" swordfish
	" table salt ",
	" tamarillo
	" tangerine
	" tap water
	" tart apple ",
	" tart crust ",
	" tater tot
	" tawny port ",
	" teff flour ",
	" thai basil ",
	" thai basils ",
	" thai chile ",
	" thai chili ",
	" thai crab
	" tiny caper ",
	" toffee bit ",
	" togarashi
	" tomatillo
	" tomato jam ",
	" tomato oil ",
	" tonka beans ",
	" topinamburs ",
	" torn basil ",
	" torn green ",
	" tortellini ",
	" trail mix
	" triple sec ",
	" trout roe
	" tuna fish
	" tuna salad ",
	" tuna steak ",
	" tuna water ",
	" turbinado
	" turkey fat ",
	" turkey ham ",
	" ugli fruits ",
	" veal bone
	" veal roast ",
	" vegenaise
	" vegetable
	" vermicelli ",
	" vin santo
	" walnut oil ",
	" watercress ",
	" watercresss ",
	" watermelon ",
	" watermelons ",
	" weisswurst ",
	" wheat beer ",
	" wheat bran ",
	" wheat bun
	" wheat germ ",
	" wheat pita ",
	" white bean ",
	" white beer ",
	" white corn ",
	" white fish ",
	" white grit ",
	" white karo ",
	" white miso ",
	" white port ",
	" white rice ",
	" white rum
	" white tuna ",
	" white wine ",
	" whitefish
	" wild boar
	" wild duck
	" wild rice
	" wood chip
	" yukon gold ",
	" yuzu juice ",
	" zinfandel
	" ziti penne ",
	" absinthe
	" advocaat
	" agar agar ",
	" aged rum
	" ahi tuna
	" allspice
	" amaranth
	" amaretti
	" amaretto
	" amber rum ",
	" andouille ",
	" anise oil ",
	" anisette
	" apple jam ",
	" apple pie ",
	" applejack ",
	" armagnac
	" arrowroot ",
	" artichoke ",
	" artichokes ",
	" asafetida ",
	" asafoetida ",
	" asparagus ",
	" asparaguss ",
	" assembly
	" aubergine ",
	" aubergines ",
	" azuki bean ",
	" baby beet ",
	" baby clam ",
	" baby corn ",
	" baby kale ",
	" baby lamb ",
	" baby leek ",
	" baby pea
	" bacon bit ",
	" bacon fat ",
	" baguette
	" baked ham ",
	" balsamic
	" barbecue
	" barberry
	" basil oil ",
	" bay leaf
	" bbq sauce ",
	" bechamel
	" beef bone ",
	" beef fat
	" beef rib
	" beef stew ",
	" beef suet ",
	" beefsteak ",
	" beetroot
	" biscotti
	" bisquick
	" black bean ",
	" blackberry ",
	" blackfish ",
	" blue crab ",
	" blueberry ",
	" blueberrys ",
	" bluefish
	" bok choy
	" bolognese ",
	" bosc pear ",
	" bouillon
	" branzino
	" bratwurst ",
	" breadfruit ",
	" breakfast ",
	" bresaola
	" broccoli
	" bucatini
	" buckwheat ",
	" budweiser ",
	" burratum
	" butternut ",
	" cacao nib ",
	" cake mix
	" cake tofu ",
	" calabreses ",
	" calamari
	" calamatum ",
	" callaloo
	" camembert ",
	" candlenut ",
	" candy bar ",
	" cantaloupe ",
	" capellini ",
	" capicola
	" capocollo ",
	" caponatum ",
	" capsicum
	" carambola ",
	" cardamaro ",
	" cardamom
	" caribbean ",
	" cavatappi ",
	" cavatelli ",
	" celeriac
	" chambord
	" chamomile ",
	" chamomiles ",
	" champagne ",
	" char siu
	" cherimoyas ",
	" chermoula ",
	" cherries
	" chestnut
	" chick pea ",
	" chickpea
	" chickpeas ",
	" chile oil ",
	" chili oil ",
	" chipotle
	" chive plu ",
	" chocolate ",
	" chopstick ",
	" chow mein ",
	" choy sum
	" ciabattum ",
	" cilantro
	" cinnamon
	" clementine ",
	" cloudberry ",
	" club soda ",
	" coca cola ",
	" cocktail
	" cocoa nib ",
	" cod steak ",
	" cointreau ",
	" coleslaw
	" condiment ",
	" cool whip ",
	" coriander ",
	" corn chip ",
	" corn cob
	" corn grit ",
	" corn husk ",
	" corn meal ",
	" corn oil
	" corn salad ",
	" cornbread ",
	" cornflake ",
	" cornflour ",
	" cornichon ",
	" cornmeal
	" courgette ",
	" courgettes ",
	" crab cake ",
	" crab meat ",
	" crabmeat
	" cranberry ",
	" cranberrys ",
	" crawfish
	" croissant ",
	" crostini
	" cucumber
	" culantro
	" cured ham ",
	" curry leaf ",
	" curry oil ",
	" dandelion ",
	" deli ham
	" deli meat ",
	" demerara
	" dill seed ",
	" dill seeds ",
	" dill weed ",
	" ditalini
	" dr pepper ",
	" drambuie
	" dressing
	" dried lime ",
	" dubonnet
	" duck egg
	" duck fat
	" duckling
	" dumpling
	" ear corn
	" egg roll
	" egg salad ",
	" egg wash
	" egg white ",
	" egg yolk
	" eggplant
	" elderberry ",
	" emmental
	" emmenthal ",
	" empanada
	" enchilada ",
	" entrecote ",
	" escabeche ",
	" escarole
	" espresso
	" everclear ",
	" falernum
	" farfalle
	" fava bean ",
	" fenugreek ",
	" fenugreeks ",
	" fettucine ",
	" fiddlehead ",
	" field pea ",
	" fingerroot ",
	" fish bone ",
	" fish cake ",
	" fish oil
	" flatbread ",
	" flavoring ",
	" flax meal ",
	" flax seed ",
	" flaxseed
	" flounder
	" focaccium ",
	" foie gra
	" framboise ",
	" fried egg ",
	" frittatum ",
	" frosting
	" fructose
	" fruit jam ",
	" furikake
	" galangal
	" gelatine
	" gem squash ",
	" gianduja
	" goat meat ",
	" goat milk ",
	" gochugaru ",
	" gochujang ",
	" goji berry ",
	" gold rum
	" goose fat ",
	" gooseberry ",
	" grapefruit ",
	" grapeseed ",
	" green bean ",
	" green pea ",
	" grenadine ",
	" guacamole ",
	" guajillo
	" guanciale ",
	" guar gum
	" guinness
	" habanero
	" half half ",
	" halloumi
	" ham bone
	" ham hock
	" ham steak ",
	" hamburger ",
	" hazelnut
	" hazelnuts ",
	" heath bar ",
	" hemp seed ",
	" herb salt ",
	" herbsaint ",
	" hero roll ",
	" hoja santa ",
	" holy basil ",
	" honeycomb ",
	" honeydew
	" hot sauce ",
	" irm tofu
	" jackfruit ",
	" jackfruits ",
	" jalapeno
	" jambalaya ",
	" jumbo egg ",
	" key lime
	" kielbasa
	" kinh giois ",
	" kiwi fruit ",
	" kiwifruit ",
	" kochujang ",
	" kochukaru ",
	" kohlrabi
	" kool aid
	" lamb bone ",
	" lamb meat ",
	" lamb rib
	" lamb stew ",
	" lavender
	" lecithin
	" lemon balm ",
	" lemon ze
	" lemonade
	" lima bean ",
	" lima beans ",
	" lime half ",
	" lime peel ",
	" linguine
	" linguini
	" liquorices ",
	" lollipop
	" long bean ",
	" macaroni
	" macaroon
	" mackerel
	" mahi mahi ",
	" mahimahi
	" manchego
	" mandarin
	" mandarines ",
	" mangetouts ",
	" mango rum ",
	" manicotti ",
	" margarine ",
	" marinade
	" marjoram
	" marmalade ",
	" marzipan
	" meat cut
	" meat loaf ",
	" meatball
	" meatloaf
	" medallion ",
	" membrillo ",
	" meringue
	" mirepoix
	" mirliton
	" mixed nut ",
	" molasses
	" monkfish
	" mulberry
	" mung bean ",
	" mung beans ",
	" mung dal
	" mushroom
	" mushrooms ",
	" navy bean ",
	" navy beans ",
	" nectarine ",
	" nectarines ",
	" newspaper ",
	" nopalito
	" nuoc cham ",
	" nuoc nam
	" nut milk
	" oat bran
	" oat flour ",
	" oil spray ",
	" okra pod
	" olive oil ",
	" oz bacon
	" oz nduja
	" paccheri
	" pak choi
	" palm oil
	" pan juice ",
	" pan spray ",
	" pancettum ",
	" panettone ",
	" parchment ",
	" parma ham ",
	" parmesan
	" partridge ",
	" passatum
	" pastrami
	" patty pans ",
	" peach jam ",
	" pecan nut ",
	" pecorino
	" pepperoni ",
	" persimmon ",
	" persimmons ",
	" pheasant
	" pie crust ",
	" pie dough ",
	" pie shell ",
	" pig tail
	" piment n
	" pimenton
	" pimiento
	" pine nut
	" pine nuts ",
	" pineapple ",
	" pineapples ",
	" pinto bean ",
	" piri piri ",
	" pistachio ",
	" plantain
	" plum jam
	" plum wine ",
	" pole bean ",
	" polentum
	" poppy seed ",
	" popsicle
	" pork butt ",
	" pork fat
	" pork lard ",
	" pork meat ",
	" pork rib
	" port wine ",
	" pot roast ",
	" potatoes
	" preserve
	" prosecco
	" provolone ",
	" purslane
	" quail egg ",
	" quick oat ",
	" raclette
	" radiatore ",
	" radicchio ",
	" radicchios ",
	" raspberry ",
	" raspberrys ",
	" reblochon ",
	" red onion ",
	" red wine
	" redcurrant ",
	" remoulade ",
	" rib roast ",
	" rib steak ",
	" rice cake ",
	" rice milk ",
	" rice mix
	" rice wine ",
	" riesling
	" rigatoni
	" rock melon ",
	" rock salt ",
	" roquefort ",
	" rose wine ",
	" rosemary
	" rosewater ",
	" ruby port ",
	" rutabaga
	" rye berry ",
	" rye bread ",
	" rye flour ",
	" sablefish ",
	" safflower ",
	" salad oil ",
	" salt cod
	" salt plu
	" salt pork ",
	" sandwich
	" sangrium
	" sassafra
	" sassafrass ",
	" sauterne
	" savoiardi ",
	" scallion
	" scallions ",
	" schmaltz
	" schnitzel ",
	" semolina
	" shallots
	" shaoxing
	" shellfish ",
	" shiitake
	" short rib ",
	" shortcake ",
	" skim milk ",
	" sloppy jo ",
	" smoothie
	" snap bean ",
	" snap pea
	" snow pea
	" soft herb ",
	" soft roll ",
	" soft tofu ",
	" sou vide
	" soup mix
	" sour mix
	" sourdough ",
	" soy cream ",
	" soy flour ",
	" soy milk
	" soy sauce ",
	" soya milk ",
	" spaetzle
	" spaghetti ",
	" spare rib ",
	" sparerib
	" spearmint ",
	" speculoo
	" spice mix ",
	" spice rub ",
	" spirulina ",
	" split pea ",
	" squid ink ",
	" squirrel
	" sriracha
	" star anise ",
	" star fruit ",
	" stew beef ",
	" stew meat ",
	" strawberry ",
	" stuffing
	" sturgeon
	" sub roll
	" succotash ",
	" sunchoke
	" sunflower ",
	" sweetcorn ",
	" swordfish ",
	" taleggio
	" tamarillo ",
	" tamarillos ",
	" tamarind
	" tangerine ",
	" tangerines ",
	" tap water ",
	" tapenade
	" tarragon
	" tater tot ",
	" tentacle
	" teriyaki
	" thai basil ",
	" thai crab ",
	" tilapium
	" tiny pea
	" togarashi ",
	" tomatillo ",
	" tomatoes
	" tonka bean ",
	" topinambur ",
	" tortilla
	" trail mix ",
	" trout roe ",
	" tuna fish ",
	" turbinado ",
	" turmeric
	" tzatziki
	" ugli fruit ",
	" umeboshi
	" urad dal
	" veal bone ",
	" vegemite
	" vegenaise ",
	" vegetable ",
	" verjuice
	" vermouth
	" vidalium
	" vin santo ",
	" watercress ",
	" watermelon ",
	" wheat bun ",
	" white rum ",
	" whitefish ",
	" wild boar ",
	" wild duck ",
	" wild rice ",
	" wood chip ",
	" zinfandel ",
	" zucchini
	" absinthe ",
	" acacium
	" advocaat ",
	" aged rum ",
	" ahi tuna ",
	" alcohol
	" alfalfa
	" allspice ",
	" allspices ",
	" almonds
	" amaranth ",
	" amaranths ",
	" amaretti ",
	" amaretto ",
	" anchovy
	" angelicas ",
	" aniseed
	" anisette ",
	" annatto
	" apricot
	" aquavit
	" armagnac ",
	" artichoke ",
	" arugula
	" asparagus ",
	" assembly ",
	" aubergine ",
	" avocado
	" baby pea ",
	" baguette ",
	" baharat
	" balsamic ",
	" bananas
	" barbecue ",
	" barberry ",
	" bay leaf ",
	" bay leafs ",
	" bechamel ",
	" beef fat ",
	" beef rib ",
	" beetroot ",
	" berbere
	" bilberrys ",
	" biscotti ",
	" biscuit
	" bisquick ",
	" blueberry ",
	" bluefish ",
	" bok choy ",
	" bok choys ",
	" bologna
	" bouillon ",
	" bourbon
	" boursin
	" branzino ",
	" bresaola ",
	" brioche
	" brisket
	" brittle
	" broccoli ",
	" broccolis ",
	" bucatini ",
	" buffalo
	" bulgogi
	" burdock
	" burratum ",
	" burrito
	" cabbage
	" cachaca
	" cake mix ",
	" calabrese ",
	" calamari ",
	" callaloo ",
	" calvado
	" campari
	" capicola ",
	" capsicum ",
	" caramel
	" caraway
	" cardamom ",
	" cardamoms ",
	" cardoon
	" carnita
	" carrots
	" cashews
	" cassava
	" cassium
	" catfish
	" cayenne
	" celeriac ",
	" celeriacs ",
	" challah
	" chambord ",
	" chamomile ",
	" char siu ",
	" chayote
	" cheddar
	" cherimoya ",
	" cherries ",
	" chervil
	" chestnut ",
	" chianti
	" chicken
	" chickpea ",
	" chickpeas ",
	" chicory
	" chipotle ",
	" chorizo
	" chowder
	" choy sum ",
	" chutney
	" cilantro ",
	" cilantros ",
	" cinnamon ",
	" cinnamons ",
	" cobbler
	" cocktail ",
	" coconut
	" codfish
	" coleslaw ",
	" collard
	" compote
	" cordial
	" corn cob ",
	" corn oil ",
	" cornmeal ",
	" courgette ",
	" couscou
	" crabmeat ",
	" cracker
	" craisin
	" cranberry ",
	" crawfish ",
	" crimini
	" crostini ",
	" crouton
	" cucumber ",
	" cucumbers ",
	" cucuzza
	" culantro ",
	" culantros ",
	" cupcake
	" cura ao
	" curacao
	" currant
	" curtido
	" custard
	" deli ham ",
	" delicatas ",
	" demerara ",
	" dessert
	" dill seed ",
	" ditalini ",
	" dogfish
	" drambuie ",
	" dressing ",
	" dubonnet ",
	" duck egg ",
	" duck fat ",
	" duckling ",
	" dumpling ",
	" ear corn ",
	" edamame
	" egg roll ",
	" egg wash ",
	" egg yolk ",
	" eggplant ",
	" eggplants ",
	" emmental ",
	" empanada ",
	" epazote
	" escarole ",
	" espresso ",
	" essence
	" falafel
	" falernum ",
	" farfalle ",
	" fatback
	" fenugreek ",
	" fig jam
	" fish oil ",
	" flaxseed ",
	" flounder ",
	" foie gra ",
	" fondant
	" fontina
	" freekeh
	" fregola
	" fritter
	" frosting ",
	" fructose ",
	" furikake ",
	" fusilli
	" gai lan
	" galanga
	" galangal ",
	" ganache
	" garnish
	" gelatin
	" gelatine ",
	" gemelli
	" genever
	" gianduja ",
	" ginseng
	" glucose
	" gnocchi
	" gold rum ",
	" granola
	" grouper
	" gruyere
	" guajillo ",
	" guar gum ",
	" guinness ",
	" habanero ",
	" habaneros ",
	" haddock
	" halibut
	" halloumi ",
	" ham bone ",
	" ham hock ",
	" harissa
	" hazelnut ",
	" hibiscu
	" honeydew ",
	" honeydews ",
	" iceberg
	" irm tofu ",
	" jackfruit ",
	" jaggery
	" jalapeno ",
	" jalapenos ",
	" jasmine
	" ketchup
	" key lime ",
	" kielbasa ",
	" kinh gioi ",
	" kohlrabi ",
	" kohlrabis ",
	" kool aid ",
	" korarimas ",
	" kumquat
	" lamb rib ",
	" lasagna
	" lasagne
	" lavender ",
	" lavenders ",
	" lecithin ",
	" lemon ze ",
	" lemonade ",
	" lettuce
	" lima bean ",
	" limeade
	" linguine ",
	" linguini ",
	" liqueur
	" liquorice ",
	" lobster
	" lollipop ",
	" macaroni ",
	" macaroon ",
	" mackerel ",
	" madeira
	" mahimahi ",
	" manchego ",
	" mandarin ",
	" mandarine ",
	" mangetout ",
	" marinade ",
	" marjoram ",
	" marjorams ",
	" marsala
	" marzipan ",
	" meat cut ",
	" meatball ",
	" meatloaf ",
	" meringue ",
	" mesclun
	" mirepoix ",
	" mirliton ",
	" mochiko
	" molasses ",
	" monkfish ",
	" mulberry ",
	" mulberrys ",
	" mung bean ",
	" mung dal ",
	" mushroom ",
	" mushrooms ",
	" mustard
	" nam pla
	" navy bean ",
	" nectarine ",
	" ni oise
	" nicoise
	" nopalito ",
	" nuoc nam ",
	" nut milk ",
	" nutella
	" oat bran ",
	" oatmeal
	" octopus
	" okra pod ",
	" old bay
	" oranges
	" oregano
	" ostrich
	" oz bacon ",
	" oz nduja ",
	" paccheri ",
	" pak choi ",
	" palm oil ",
	" pancake
	" paprika
	" parmesan ",
	" parsley
	" parsnip
	" passatum ",
	" pastina
	" pastrami ",
	" patty pan ",
	" pea pod
	" peaches
	" peanuts
	" pecorino ",
	" persimmon ",
	" pheasant ",
	" physaliss ",
	" pierogi
	" pig tail ",
	" pignoli
	" piment n ",
	" pimento
	" pimenton ",
	" pimiento ",
	" pine nut ",
	" pineapple ",
	" pinenut
	" plantain ",
	" plum jam ",
	" poblano
	" polentum ",
	" pollock
	" pompano
	" popcorn
	" popsicle ",
	" porcini
	" pork fat ",
	" pork rib ",
	" potatoes ",
	" poultry
	" praline
	" preserve ",
	" pretzel
	" prosecco ",
	" pudding
	" pumpkin
	" purslane ",
	" raclette ",
	" radicchio ",
	" raisins
	" rambutans ",
	" raspberry ",
	" ravioli
	" recaito
	" red wine ",
	" redfish
	" rhubarb
	" rice mix ",
	" riesling ",
	" rigatoni ",
	" risotto
	" ro wine
	" romaine
	" rosebud
	" rosemary ",
	" rosemarys ",
	" rouille
	" rutabaga ",
	" rutabagas ",
	" saffron
	" salt cod ",
	" salt plu ",
	" saltine
	" sambuca
	" sandwich ",
	" sangrium ",
	" sardine
	" sassafra ",
	" sassafras ",
	" sausage
	" sauterne ",
	" scallion ",
	" scallions ",
	" scallop
	" schmaltz ",
	" schnapp
	" seafood
	" seaweed
	" sel gri
	" seltzer
	" semolina ",
	" serrano
	" shallot
	" shallots ",
	" shaoxing ",
	" shiitake ",
	" sirloin
	" smoothie ",
	" snap pea ",
	" snap peas ",
	" snapper
	" snow pea ",
	" sofrito
	" sorghum
	" sou vide ",
	" souffle
	" soup mix ",
	" sour mix ",
	" soy beans ",
	" soy milk ",
	" soybean
	" soymilk
	" spaetzle ",
	" spanish
	" sparerib ",
	" speculoo ",
	" spinach
	" splenda
	" squirrel ",
	" sriracha ",
	" stevium
	" stilton
	" stuffing ",
	" sturgeon ",
	" sub roll ",
	" sucanat
	" sultana
	" sunchoke ",
	" tabasco
	" taleggio ",
	" tamarillo ",
	" tamarind ",
	" tamarinds ",
	" tangerine ",
	" tapenade ",
	" tapioca
	" tarragon ",
	" tarragons ",
	" tentacle ",
	" tequila
	" teriyaki ",
	" tilapium ",
	" tiny pea ",
	" tomatoes ",
	" tortilla ",
	" tostada
	" treacle
	" treviso
	" tri tip
	" trotter
	" truffle
	" truvium
	" turkish
	" turmeric ",
	" turmerics ",
	" tzatziki ",
	" umeboshi ",
	" urad dal ",
	" vanilla
	" vegemite ",
	" venison
	" verjuice ",
	" vermouth ",
	" vidalium ",
	" vinegar
	" walleye
	" walnuts
	" whiskey
	" whiting
	" woodruffs ",
	" yoghurt
	" z vodka
	" za atar
	" zucchini ",
	" zucchinis ",
	" acacium ",
	" add in
	" ajwain
	" alcohol ",
	" alfalfa ",
	" allspice ",
	" almond
	" almonds ",
	" amaranth ",
	" amchoors ",
	" amchur
	" anchovy ",
	" angelica ",
	" aniseed ",
	" annatto ",
	" aperol
	" apples
	" apricot ",
	" apricots ",
	" aquavit ",
	" arugula ",
	" arugulas ",
	" asiago
	" avocado ",
	" avocados ",
	" baharat ",
	" banana
	" bananas ",
	" barley
	" barolo
	" bay leaf ",
	" berbere ",
	" bilberry ",
	" biscuit ",
	" bisque
	" bok choy ",
	" bologna ",
	" bonito
	" bourbon ",
	" boursin ",
	" brandy
	" brioche ",
	" brisket ",
	" brittle ",
	" broccoli ",
	" browny
	" buffalo ",
	" bulgogi ",
	" bulgur
	" burdock ",
	" burger
	" burrito ",
	" butter
	" cabbage ",
	" cabbages ",
	" cachaca ",
	" calvado ",
	" campari ",
	" camphors ",
	" canola
	" cantal
	" capers
	" caramel ",
	" caraway ",
	" caraways ",
	" cardamom ",
	" cardoon ",
	" carnita ",
	" carrot
	" carrots ",
	" cashew
	" cashews ",
	" cassava ",
	" cassium ",
	" catfish ",
	" catsup
	" caviar
	" cayenne ",
	" celeriac ",
	" celery
	" cereal
	" ch vre
	" challah ",
	" charolis ",
	" chayote ",
	" chayotes ",
	" cheddar ",
	" cheese
	" cherry
	" chervil ",
	" chervils ",
	" chianti ",
	" chicken ",
	" chickpea ",
	" chicory ",
	" chives
	" chorizo ",
	" chowder ",
	" chutney ",
	" cilantro ",
	" cinnamon ",
	" citron
	" cobbler ",
	" cockle
	" coconut ",
	" coconuts ",
	" codfish ",
	" coffee
	" cognac
	" collard ",
	" comice
	" compote ",
	" confit
	" cookie
	" cordial ",
	" cotija
	" couscou ",
	" cracker ",
	" craisin ",
	" creamy
	" crimini ",
	" crisco
	" crouton ",
	" cucumber ",
	" cucuzza ",
	" culantro ",
	" cupcake ",
	" cura ao ",
	" curacao ",
	" currant ",
	" currants ",
	" curtido ",
	" custard ",
	" daikon
	" delicata ",
	" dessert ",
	" dogfish ",
	" dragee
	" dukkah
	" edamame ",
	" eggnog
	" eggplant ",
	" endife
	" endive
	" epazote ",
	" epazotes ",
	" essence ",
	" fajita
	" falafel ",
	" farina
	" fatback ",
	" fennel
	" fig jam ",
	" fondant ",
	" fondue
	" fontina ",
	" freekeh ",
	" fregola ",
	" french
	" fresca
	" frisee
	" fritter ",
	" fusilli ",
	" gai lan ",
	" galanga ",
	" ganache ",
	" garlic
	" garnish ",
	" gelatin ",
	" gelato
	" gemelli ",
	" genever ",
	" ginger
	" ginseng ",
	" glucose ",
	" gluten
	" gnocchi ",
	" granola ",
	" grappa
	" gratin
	" grease
	" grouper ",
	" gruyere ",
	" habanero ",
	" haddock ",
	" halibut ",
	" halvah
	" harissa ",
	" hibiscu ",
	" hijiki
	" hoisin
	" hominy
	" honeydew ",
	" iceberg ",
	" injera
	" jaggery ",
	" jalapeno ",
	" jasmine ",
	" jicama
	" kahlua
	" kaiser
	" kernel
	" ketchup ",
	" kimchi
	" kirsch
	" kohlrabi ",
	" korarima ",
	" kumara
	" kumquat ",
	" kumquats ",
	" labneh
	" lardon
	" lasagna ",
	" lasagne ",
	" laurel
	" lavash
	" lavender ",
	" lemons
	" lentil
	" lentilss ",
	" lettuce ",
	" lettuces ",
	" lillet
	" limeade ",
	" lingui
	" liqueur ",
	" liquor
	" lobster ",
	" loquat
	" lovage
	" lychee
	" madeira ",
	" maldon
	" marjoram ",
	" marsala ",
	" masala
	" mastic
	" matcha
	" matzoh
	" mesclun ",
	" mezcal
	" midori
	" millet
	" mitsubas ",
	" mizuna
	" mochiko ",
	" mousse
	" muesli
	" muffin
	" mugworts ",
	" mulberry ",
	" mushroom ",
	" mussel
	" mustard ",
	" mutton
	" myrtle
	" nam pla ",
	" nectar
	" ni oise ",
	" nicoise ",
	" nigellas ",
	" njangsas ",
	" noodle
	" nopale
	" nougat
	" noyaux
	" nutella ",
	" nutmeg
	" oatmeal ",
	" octopus ",
	" old bay ",
	" olives
	" onions
	" orange
	" oranges ",
	" orchid
	" oregano ",
	" oreganos ",
	" orgeat
	" ostrich ",
	" oxtail
	" oyster
	" pancake ",
	" paneer
	" panela
	" papaya
	" paprika ",
	" paprikas ",
	" parsley ",
	" parsleys ",
	" parsnip ",
	" parsnips ",
	" pastina ",
	" pastry
	" pawpaw
	" pea pod ",
	" peaches ",
	" peanut
	" peanuts ",
	" pecans
	" pectin
	" pepita
	" pepper
	" perillas ",
	" pernod
	" phyllo
	" physalis ",
	" pickle
	" pierogi ",
	" pignoli ",
	" pimento ",
	" pinenut ",
	" pistou
	" poblano ",
	" pollock ",
	" pomelo
	" pompano ",
	" popcorn ",
	" porcini ",
	" potato
	" poultry ",
	" praline ",
	" pretzel ",
	" pudding ",
	" pumpkin ",
	" pumpkins ",
	" quahog
	" quiche
	" quince
	" quinoa
	" rabbit
	" radhunis ",
	" radish
	" raisin
	" raisins ",
	" rambutan ",
	" rapini
	" ravioli ",
	" recaito ",
	" redfish ",
	" relish
	" rennet
	" rhubarb ",
	" rhubarbs ",
	" ribeye
	" risotto ",
	" ro wine ",
	" rocket
	" romaine ",
	" rooibo
	" rosebud ",
	" rosemary ",
	" rotini
	" rouille ",
	" russet
	" rutabaga ",
	" saffron ",
	" saffrons ",
	" salami
	" salmon
	" saltine ",
	" sambal
	" sambuca ",
	" sardine ",
	" satsumas ",
	" sausage ",
	" savory
	" scallion ",
	" scallop ",
	" scampi
	" schnapp ",
	" scotch
	" seafood ",
	" seaweed ",
	" seitan
	" sel gri ",
	" seltzer ",
	" serrano ",
	" sesame
	" shallot ",
	" shallots ",
	" sherry
	" shrimp
	" sirloin ",
	" skewer
	" skirrets ",
	" snap pea ",
	" snapper ",
	" sofrito ",
	" sorbet
	" sorghum ",
	" sorrel
	" souffle ",
	" soy bean ",
	" soybean ",
	" soymilk ",
	" spanish ",
	" spinach ",
	" spinachs ",
	" spirit
	" splenda ",
	" spread
	" sprite
	" sprout
	" squash
	" squash s ",
	" starch
	" stevium ",
	" stilton ",
	" sucanat ",
	" sultana ",
	" tabasco ",
	" tahini
	" tamale
	" tamari
	" tamarind ",
	" tapioca ",
	" tarama
	" tarragon ",
	" tat sois ",
	" tempeh
	" tequila ",
	" tobiko
	" toffee
	" tomato
	" tostada ",
	" treacle ",
	" treviso ",
	" tri tip ",
	" trotter ",
	" truffle ",
	" truvium ",
	" turbot
	" turkey
	" turkish ",
	" turmeric ",
	" turnip
	" vanilla ",
	" vanillas ",
	" venison ",
	" vinegar ",
	" waffle
	" wakame
	" walleye ",
	" walnut
	" walnuts ",
	" wasabi
	" whiskey ",
	" whisky
	" whiting ",
	" wonton
	" woodruff ",
	" yoghurt ",
	" yogurt
	" z vodka ",
	" za atar ",
	" zedoarys ",
	" zereshks ",
	" zucchini ",
	" add in ",
	" adobo
	" agave
	" aioli
	" ajwain ",
	" ajwains ",
	" almond ",
	" amaro
	" amber
	" amchoor ",
	" amchur ",
	" ancho
	" anise
	" aonoris ",
	" aperol ",
	" apple
	" apples ",
	" apricot ",
	" arame
	" arugula ",
	" asiago ",
	" avocado ",
	" bacon
	" bagel
	" banana ",
	" bananas ",
	" barley ",
	" barolo ",
	" basil
	" beans
	" beets
	" berry
	" bison
	" bisque ",
	" bonito ",
	" borages ",
	" brandy ",
	" bread
	" bream
	" brine
	" broth
	" browny ",
	" bugle
	" bulgur ",
	" burger ",
	" butter ",
	" cabbage ",
	" cacao
	" cacha
	" cactu
	" cajun
	" camphor ",
	" candy
	" canola ",
	" cantal ",
	" caper
	" capers ",
	" caraway ",
	" carob
	" carrot ",
	" carrots ",
	" cashew ",
	" cassi
	" cassias ",
	" catsup ",
	" caviar ",
	" cedar
	" celery ",
	" celerys ",
	" cereal ",
	" ch vre ",
	" chard
	" charoli ",
	" chayote ",
	" cheese ",
	" chenpis ",
	" cherry ",
	" cherrys ",
	" chervil ",
	" chile
	" chili
	" chive
	" chives ",
	" chivess ",
	" chuck
	" cicelys ",
	" cider
	" citron ",
	" citru
	" clove
	" cockle ",
	" cocoa
	" coconut ",
	" coffee ",
	" cognac ",
	" colby
	" comice ",
	" conch
	" confit ",
	" cookie ",
	" copha
	" coppa
	" cotija ",
	" cream
	" creamy ",
	" crema
	" crepe
	" crisco ",
	" crisp
	" crust
	" cumin
	" currant ",
	" curry
	" cynar
	" daikon ",
	" daikons ",
	" damsons ",
	" dashi
	" dates
	" dijon
	" donut
	" dough
	" dragee ",
	" drink
	" dukkah ",
	" durians ",
	" eggnog ",
	" endife ",
	" endive ",
	" endives ",
	" epazote ",
	" equal
	" fajita ",
	" farina ",
	" farro
	" feijoas ",
	" fennel ",
	" fennels ",
	" fetum
	" fideo
	" filet
	" flour
	" fluke
	" fondue ",
	" frank
	" french ",
	" fresca ",
	" fri e
	" frisee ",
	" frisees ",
	" frito
	" fruit
	" fudge
	" garlic ",
	" garlics ",
	" gelato ",
	" ginger ",
	" gingers ",
	" gluten ",
	" golpars ",
	" goose
	" gouda
	" grape
	" grappa ",
	" grass
	" gratin ",
	" gravy
	" grease ",
	" greek
	" green
	" guava
	" gumbo
	" halvah ",
	" herbe
	" hijiki ",
	" hoisin ",
	" hominy ",
	" honey
	" hummu
	" hyssops ",
	" icing
	" injera ",
	" jambuls ",
	" jello
	" jelly
	" jicama ",
	" jicamas ",
	" juice
	" jujubes ",
	" kabob
	" kahlua ",
	" kaiser ",
	" kamut
	" kasha
	" kebab
	" kefir
	" kernel ",
	" kimchi ",
	" kirsch ",
	" kombu
	" konbu
	" kumara ",
	" kumquat ",
	" labneh ",
	" lager
	" lardo
	" lardon ",
	" laurel ",
	" lavash ",
	" leeks
	" legumes ",
	" lemon
	" lemons ",
	" lentil ",
	" lentils ",
	" lettuce ",
	" licor
	" lillet ",
	" lingui ",
	" liquor ",
	" liver
	" loafe
	" loquat ",
	" loquats ",
	" lovage ",
	" lovages ",
	" lychee ",
	" lychees ",
	" m che
	" mahlabs ",
	" maldon ",
	" mango
	" maple
	" masala ",
	" mastic ",
	" mastics ",
	" matcha ",
	" matzo
	" matzoh ",
	" melon
	" mezcal ",
	" midori ",
	" millet ",
	" mirin
	" mitsuba ",
	" mizuna ",
	" moose
	" morel
	" mousse ",
	" muesli ",
	" muffin ",
	" mugwort ",
	" mussel ",
	" mutton ",
	" myrtle ",
	" nacho
	" nduja
	" nectar ",
	" nigella ",
	" njangsa ",
	" noodle ",
	" nopal
	" nopale ",
	" nopales ",
	" nougat ",
	" noyaux ",
	" nutmeg ",
	" nutmegs ",
	" olive
	" olives ",
	" onion
	" onions ",
	" orange ",
	" oranges ",
	" orchid ",
	" oregano ",
	" orgeat ",
	" oxtail ",
	" oyster ",
	" pamelos ",
	" paneer ",
	" panela ",
	" panir
	" panko
	" pansy
	" papaya ",
	" papayas ",
	" paprika ",
	" parsley ",
	" parsnip ",
	" pasta
	" pasti
	" pastry ",
	" pawpaw ",
	" peach
	" peanut ",
	" pears
	" pecan
	" pecans ",
	" pectin ",
	" penne
	" pepita ",
	" pepper ",
	" perch
	" perilla ",
	" pernod ",
	" pesto
	" phyllo ",
	" pickle ",
	" pilaf
	" pisco
	" pistou ",
	" pitum
	" pizza
	" pluot
	" pomelo ",
	" pomelos ",
	" ponzu
	" potato ",
	" potatos ",
	" prawn
	" prune
	" pumpkin ",
	" quahog ",
	" quail
	" quark
	" queso
	" quiche ",
	" quince ",
	" quinces ",
	" quinoa ",
	" rabbit ",
	" radhuni ",
	" radish ",
	" radishs ",
	" raisin ",
	" raisins ",
	" raman
	" rapini ",
	" relish ",
	" rennet ",
	" rhubarb ",
	" ribeye ",
	" roast
	" rocket ",
	" rooibo ",
	" rotel
	" rotini ",
	" russet ",
	" saffron ",
	" salad
	" salami ",
	" salmon ",
	" salsa
	" sambal ",
	" sanshos ",
	" satsuma ",
	" sauce
	" savory ",
	" savorys ",
	" scampi ",
	" scone
	" scotch ",
	" scrod
	" seitan ",
	" sesame ",
	" sesames ",
	" shallot ",
	" shark
	" sheep
	" sherry ",
	" shiso
	" shoyu
	" shrimp ",
	" skewer ",
	" skirret ",
	" smelt
	" snail
	" sorbet ",
	" sorrel ",
	" sorrels ",
	" speck
	" spelt
	" spinach ",
	" spirit ",
	" spray
	" spread ",
	" sprite ",
	" sprout ",
	" squash ",
	" squash ",
	" squid
	" starch ",
	" steak
	" stout
	" sugar
	" sumac
	" sumaq
	" swede
	" syrup
	" tahini ",
	" tamale ",
	" tamari ",
	" tarama ",
	" tasso
	" tat soi ",
	" tempeh ",
	" thyme
	" toast
	" tobiko ",
	" toffee ",
	" tomato ",
	" torte
	" tripe
	" trout
	" tuberss ",
	" turbot ",
	" turkey ",
	" turnip ",
	" turnips ",
	" vanilla ",
	" verju
	" vodka
	" wafer
	" waffle ",
	" wakame ",
	" walnut ",
	" wasabi ",
	" wasabis ",
	" water
	" wheat
	" whisky ",
	" wonton ",
	" yeast
	" yogurt ",
	" yucca
	" zedoary ",
	" zereshk ",
	" adobo ",
	" agar
	" agave ",
	" aioli ",
	" ajwain ",
	" alum
	" amaro ",
	" amber ",
	" ancho ",
	" anise ",
	" anises ",
	" aonori ",
	" apple ",
	" apples ",
	" arame ",
	" bacon ",
	" bagel ",
	" banana ",
	" bark
	" basil ",
	" basils ",
	" bass
	" bean
	" beans ",
	" beef
	" beer
	" beet
	" beets ",
	" berry ",
	" bing
	" bison ",
	" boar
	" boldos ",
	" bone
	" borage ",
	" bran
	" bread ",
	" bream ",
	" brie
	" brine ",
	" broth ",
	" bugle ",
	" cacao ",
	" cacha ",
	" cactu ",
	" cajun ",
	" cake
	" candy ",
	" caper ",
	" carob ",
	" carp
	" carrot ",
	" cassi ",
	" cassia ",
	" cava
	" cedar ",
	" celery ",
	" chai
	" chard ",
	" chards ",
	" chee
	" chenpi ",
	" cherry ",
	" chile ",
	" chili ",
	" chip
	" chive ",
	" chives ",
	" chuck ",
	" cicely ",
	" cider ",
	" citru ",
	" clam
	" clov
	" clove ",
	" cloves ",
	" cocoa ",
	" cola
	" colby ",
	" conch ",
	" copha ",
	" coppa ",
	" corn
	" cr m
	" crab
	" cream ",
	" crema ",
	" crepe ",
	" cresss ",
	" crisp ",
	" crust ",
	" cubebs ",
	" cumin ",
	" cumins ",
	" curd
	" curry ",
	" cynar ",
	" daikon ",
	" damson ",
	" dashi ",
	" date
	" dates ",
	" dijon ",
	" dill
	" donut ",
	" dough ",
	" drink ",
	" duck
	" durian ",
	" edam
	" eggs
	" endive ",
	" equal ",
	" farro ",
	" feijoa ",
	" fennel ",
	" fetum ",
	" fideo ",
	" filet ",
	" fish
	" flax
	" flour ",
	" fluke ",
	" frank ",
	" fri e ",
	" frisee ",
	" frito ",
	" fruit ",
	" fudge ",
	" garlic ",
	" ghee
	" ginger ",
	" goat
	" gold
	" golpar ",
	" goose ",
	" gouda ",
	" grape ",
	" grapes ",
	" grass ",
	" gravy ",
	" greek ",
	" green ",
	" grit
	" gruy
	" guava ",
	" guavas ",
	" gumbo ",
	" hake
	" hash
	" herb
	" herbe ",
	" hing
	" honey ",
	" hummu ",
	" hyssop ",
	" icing ",
	" jambul ",
	" jello ",
	" jelly ",
	" jerk
	" jicama ",
	" jimbus ",
	" juice ",
	" jujube ",
	" kabob ",
	" kahl
	" kale
	" kamut ",
	" kasha ",
	" kebab ",
	" kefir ",
	" kelp
	" kiwi
	" kokums ",
	" kombu ",
	" konbu ",
	" lager ",
	" lamb
	" lard
	" lardo ",
	" leek
	" leek s ",
	" leeks ",
	" legume ",
	" lemon ",
	" lemons ",
	" licor ",
	" lime
	" liver ",
	" loafe ",
	" loquat ",
	" lotu
	" lovage ",
	" lychee ",
	" m che ",
	" mace
	" mahlab ",
	" malt
	" mameys ",
	" mango ",
	" mangos ",
	" maple ",
	" masa
	" mastic ",
	" mate
	" matzo ",
	" meat
	" melon ",
	" milk
	" mint
	" mirin ",
	" miso
	" moose ",
	" morel ",
	" naan
	" nacho ",
	" nduja ",
	" nopal ",
	" nopale ",
	" nori
	" nutmeg ",
	" nuts
	" oats
	" okra
	" oleo
	" olive ",
	" olives ",
	" onion ",
	" onions ",
	" orange ",
	" oreo
	" orzo
	" ouzo
	" palm
	" pamelo ",
	" panir ",
	" panko ",
	" pansy ",
	" papaya ",
	" pasta ",
	" pasti ",
	" pate
	" peach ",
	" peachs ",
	" pear
	" pears ",
	" peas
	" pecan ",
	" penne ",
	" perch ",
	" pesto ",
	" pilaf ",
	" pimm
	" pine
	" pisco ",
	" pita
	" pitum ",
	" pizza ",
	" plum
	" pluot ",
	" pomelo ",
	" ponzu ",
	" pork
	" port
	" potato ",
	" prawn ",
	" prune ",
	" quail ",
	" quark ",
	" queso ",
	" quince ",
	" radish ",
	" ragu
	" raisin ",
	" raman ",
	" ramp
	" rice
	" roast ",
	" rose
	" rotel ",
	" roti
	" rump
	" saba
	" sage
	" sago
	" sake
	" salad ",
	" salsa ",
	" salt
	" sansho ",
	" sauce ",
	" savory ",
	" scone ",
	" scrod ",
	" seed
	" sesame ",
	" shark ",
	" sheep ",
	" shiso ",
	" shisos ",
	" shoyu ",
	" skin
	" smelt ",
	" snail ",
	" soba
	" soda
	" soju
	" sole
	" sorrel ",
	" soup
	" spam
	" speck ",
	" spelt ",
	" spray ",
	" squid ",
	" steak ",
	" stew
	" stout ",
	" suet
	" sugar ",
	" sumac ",
	" sumacs ",
	" sumaq ",
	" swede ",
	" syrup ",
	" taco
	" taro
	" tart
	" tasso ",
	" teff
	" thru
	" thyme ",
	" thymes ",
	" toast ",
	" tofu
	" torte ",
	" tripe ",
	" trout ",
	" tubers ",
	" tuna
	" turnip ",
	" udon
	" urfa
	" uzazis ",
	" veal
	" verju ",
	" vodka ",
	" wafer ",
	" wasabi ",
	" water ",
	" wheat ",
	" whey
	" wine
	" wrap
	" yeast ",
	" yolk
	" yuca
	" yucca ",
	" ziti
	" agar ",
	" ahi
	" ale
	" alum ",
	" ame
	" anise ",
	" apple ",
	" bark ",
	" basil ",
	" bass ",
	" bay
	" bean ",
	" beef ",
	" beer ",
	" beet ",
	" beets ",
	" bing ",
	" boar ",
	" boldo ",
	" bone ",
	" bran ",
	" brie ",
	" bun
	" cake ",
	" cal
	" carp ",
	" cava ",
	" chai ",
	" chard ",
	" chee ",
	" chip ",
	" clam ",
	" clov ",
	" clove ",
	" cod
	" cola ",
	" corn ",
	" corns ",
	" cr m ",
	" crab ",
	" cress ",
	" cubeb ",
	" cumin ",
	" curd ",
	" date ",
	" dates ",
	" dill ",
	" dills ",
	" duck ",
	" edam ",
	" eel
	" egg
	" eggs ",
	" fat
	" fig
	" fish ",
	" flax ",
	" ghee ",
	" gin
	" goat ",
	" gold ",
	" grape ",
	" grit ",
	" gruy ",
	" guava ",
	" hake ",
	" ham
	" hash ",
	" hemps ",
	" herb ",
	" hing ",
	" hop
	" ice
	" jam
	" jerk ",
	" jimbu ",
	" kahl ",
	" kale ",
	" kales ",
	" kelp ",
	" kiwi ",
	" kokum ",
	" lamb ",
	" lard ",
	" leek ",
	" leek ",
	" lemon ",
	" lime ",
	" limes ",
	" lotu ",
	" lox
	" m m
	" mace ",
	" maces ",
	" malt ",
	" mamey ",
	" mango ",
	" masa ",
	" mate ",
	" meat ",
	" milk ",
	" mint ",
	" mints ",
	" miso ",
	" msg
	" naan ",
	" nori ",
	" nut
	" nuts ",
	" oat
	" oats ",
	" oil
	" okra ",
	" okras ",
	" oleo ",
	" olive ",
	" onion ",
	" oreo ",
	" orzo ",
	" ouzo ",
	" palm ",
	" pate ",
	" pea
	" peach ",
	" pear ",
	" pears ",
	" peas ",
	" pie
	" pimm ",
	" pine ",
	" pita ",
	" plu
	" plum ",
	" plums ",
	" pork ",
	" port ",
	" ragu ",
	" ramp ",
	" rib
	" rice ",
	" roe
	" rose ",
	" roses ",
	" roti ",
	" rub
	" rum
	" rump ",
	" rye
	" saba ",
	" sage ",
	" sages ",
	" sago ",
	" sake ",
	" salt ",
	" salts ",
	" seed ",
	" shiso ",
	" skin ",
	" soba ",
	" soda ",
	" soju ",
	" sole ",
	" soup ",
	" spam ",
	" stew ",
	" suet ",
	" sumac ",
	" taco ",
	" taro ",
	" taros ",
	" tart ",
	" teff ",
	" thru ",
	" thyme ",
	" tip
	" tofu ",
	" tuna ",
	" udon ",
	" urfa ",
	" uzazi ",
	" veal ",
	" whey ",
	" wine ",
	" wrap ",
	" yam
	" yolk ",
	" yuca ",
	" yuzus ",
	" zests ",
	" ziti ",
	" ahi ",
	" ale ",
	" ame ",
	" bay ",
	" beet ",
	" bun ",
	" cal ",
	" cod ",
	" corn ",
	" date ",
	" dill ",
	" eel ",
	" egg ",
	" fat ",
	" fig ",
	" figs ",
	" gin ",
	" ham ",
	" hemp ",
	" hop ",
	" ice ",
	" jam ",
	" kale ",
	" lime ",
	" lox ",
	" m m ",
	" mace ",
	" mint ",
	" ml
	" msg ",
	" nut ",
	" nuts ",
	" oat ",
	" oil ",
	" okra ",
	" pea ",
	" pear ",
	" peas ",
	" pie ",
	" plu ",
	" plum ",
	" rib ",
	" ro
	" roe ",
	" rose ",
	" rub ",
	" rues ",
	" rum ",
	" rye ",
	" sage ",
	" salt ",
	" taro ",
	" tip ",
	" yam ",
	" yams ",
	" yuzu ",
	" zest ",
	" fig ",
	" ml ",
	" nut ",
	" pea ",
	" ro ",
	" rue ",
	" yam "}
var corpusMeasures = []string{" tablespoons ",
	" milliliter ",
	" tablespoon ",
	" teaspoons ",
	" teaspoon ",
	" canned ",
	" ounces ",
	" pounds ",
	" quarts ",
	" grams ",
	" ounce ",
	" pints ",
	" pound ",
	" quart ",
	" tbsps ",
	" cans ",
	" cups ",
	" gram ",
	" pint ",
	" tbsp ",
	" tsps ",
	" can ",
	" cup ",
	" tbl ",
	" tsp ",
	" ml ",
	" oz ",
	" c ",
	" g "}
var corpusNumbers = []string{" 1/2 ",
	" 1/3 ",
	" 1/4 ",
	" 1/5 ",
	" 1/6 ",
	" 1/7 ",
	" 1/8 ",
	" 2/3 ",
	" 2/5 ",
	" 2/7 ",
	" 3/4 ",
	" 3/8 ",
	" 4/5 ",
	" 5/8 ",
	" 7/8 ",
	" 1 ",
	" 2 ",
	" 3 ",
	" 4 ",
	" 5 ",
	" 6 ",
	" 7 ",
	" 8 ",
	" 9 ",
	" 10 ",
	" 11 ",
	" 12 ",
	" 13 ",
	" 14 ",
	" 15 ",
	" 16 ",
	" 17 ",
	" 18 ",
	" 19 ",
	" 20 ",
	" ⅜ ",
	" ⅞ ",
	" ⅔ ",
	" ⅝ ",
	" ⅓ ",
	" ½ ",
	" ¼ ",
	" ¾ ",
	" ⅛ "}
var corpusDirections = []string{" 1 ",
	" 10 ",
	" 15 ",
	" 2 ",
	" 25 ",
	" 375 ",
	" 6 ",
	" 8 ",
	" 9x13 ",
	" a ",
	" about ",
	" achieve ",
	" add ",
	" additional ",
	" an ",
	" and ",
	" arrange ",
	" assemble ",
	" bake ",
	" baking ",
	" basil ",
	" beef ",
	" before ",
	" boil ",
	" boiling ",
	" bottom ",
	" bowl ",
	" bring ",
	" brown ",
	" browned ",
	" butter ",
	" cheese ",
	" chopped ",
	" coconut ",
	" cold ",
	" color ",
	" combine ",
	" completely ",
	" cook ",
	" cookie ",
	" cooking ",
	" cool ",
	" corn ",
	" cover ",
	" covered ",
	" crushed ",
	" cup ",
	" cups ",
	" degrees ",
	" dish ",
	" distributed ",
	" does ",
	" drain ",
	" dutch ",
	" eggs ",
	" either ",
	" even ",
	" evenly ",
	" everything ",
	" f ",
	" fennel ",
	" foil ",
	" for ",
	" from ",
	" garlic ",
	" grab ",
	" ground ",
	" half ",
	" heat ",
	" hours ",
	" in ",
	" inch ",
	" into ",
	" is ",
	" italian ",
	" large ",
	" lasagna ",
	" layers ",
	" lengthwise ",
	" let ",
	" lightly ",
	" make ",
	" meat ",
	" medium ",
	" melted ",
	" minutes ",
	" mix ",
	" mixing ",
	" mixture ",
	" mozzarella ",
	" next ",
	" noodles ",
	" not ",
	" nutmeg ",
	" nuts ",
	" oats ",
	" occasionally ",
	" of ",
	" on ",
	" one ",
	" onion ",
	" or ",
	" oven ",
	" over ",
	" pans ",
	" parmesan ",
	" parsley ",
	" paste ",
	" pecans ",
	" pepper ",
	" pie ",
	" place ",
	" pot ",
	" pour ",
	" preheat ",
	" preheated ",
	" prevent ",
	" put ",
	" raisins ",
	" remaining ",
	" remove ",
	" repeat ",
	" ricotta ",
	" rinse ",
	" salt ",
	" salted ",
	" sauce ",
	" sausage ",
	" season ",
	" seasoning ",
	" seeds ",
	" serving ",
	" sheet ",
	" shell ",
	" simmer ",
	" slices ",
	" spoon ",
	" spray ",
	" spread ",
	" sprinkle ",
	" sticking ",
	" stir ",
	" stirring ",
	" sugar ",
	" sure ",
	" syrup ",
	" tablespoon ",
	" tablespoons ",
	" teaspoon ",
	" the ",
	" then ",
	" third ",
	" to ",
	" tomato ",
	" tomatoes ",
	" top ",
	" touch ",
	" transfer ",
	" until ",
	" water ",
	" well ",
	" with "}
var corpusDirectionsNeg = []string{" a ",
	" all ",
	" also ",
	" amount ",
	" and ",
	" be ",
	" best ",
	" brand ",
	" butter ",
	" c ",
	" calcium ",
	" calories ",
	" carbohydrates ",
	" chip ",
	" chocolate ",
	" cholesterol ",
	" cookie ",
	" cookies. ",
	" costco ",
	" dessert ",
	" dinner ",
	" dough ",
	" ensure ",
	" equally ",
	" ever ",
	" excellent ",
	" fat ",
	" fiber ",
	" for ",
	" from ",
	" great! ",
	" have ",
	" i ",
	" images ",
	" iron ",
	" it's ",
	" just ",
	" kirkland ",
	" liking! ",
	" make ",
	" per ",
	" potassium ",
	" protein ",
	" recipe ",
	" recommend ",
	" results. ",
	" salted ",
	" saturated ",
	" serving ",
	" servings ",
	" sodium ",
	" tasting ",
	" text ",
	" that ",
	" the ",
	" then ",
	" these ",
	" tillamook ",
	" to ",
	" unsalted ",
	" use ",
	" used ",
	" vitamin ",
	" with ",
	" would ",
	" yield ",
	" your "}

type fractionNumber struct {
	fractionString string
	value          float64
}

var corpusFractionNumberMap = map[string]fractionNumber{
	"¼": {"1/4", 0.2500000000},
	"½": {"1/2", 0.5000000000},
	"¾": {"3/4", 0.7500000000},
	"⅓": {"1/3", 0.3333333333},
	"⅔": {"2/3", 0.6666666667},
	"⅛": {"1/8", 0.1250000000},
	"⅜": {"3/8", 0.3750000000},
	"⅝": {"5/8", 0.6250000000},
	"⅞": {"7/8", 0.8750000000},
}

var corpusMeasuresMap = map[string]string{
	"c":           "cup",
	"can":         "can",
	"canned":      "can",
	"cans":        "can",
	"cup":         "cup",
	"cups":        "cup",
	"g":           "gram",
	"gram":        "gram",
	"grams":       "gram",
	"milliliter":  "milliliter",
	"ml":          "milliliter",
	"ounce":       "ounce",
	"ounces":      "ounce",
	"oz":          "ounce",
	"pint":        "pint",
	"pints":       "pint",
	"pound":       "pound",
	"pounds":      "pound",
	"quart":       "quart",
	"quarts":      "quart",
	"tablespoon":  "tbl",
	"tablespoons": "tbl",
	"tbl":         "tbl",
	"tbsp":        "tbl",
	"tbsps":       "tbl",
	"teaspoon":    "tsp",
	"teaspoons":   "tsp",
	"tsp":         "tsp",
	"tsps":        "tsp",
}

var densities = map[string]float64{
	"almond milk":            245.5000000000,
	"almonds":                115.1000000000,
	"apple":                  144.6000000000,
	"apple juice":            243.5000000000,
	"apples":                 152.2000000000,
	"applesauce":             247.2000000000,
	"arugula":                20.0000000000,
	"asparagus":              204.3000000000,
	"avocado":                218.0000000000,
	"bacon":                  156.8000000000,
	"balsamic vinegar":       255.0000000000,
	"banana":                 159.1000000000,
	"bananas":                158.3000000000,
	"barbecue sauce":         279.2000000000,
	"basil":                  191.0000000000,
	"beans":                  189.5000000000,
	"beef":                   212.8000000000,
	"beef broth":             225.2000000000,
	"beef stock":             245.3000000000,
	"beer":                   237.0000000000,
	"beets":                  195.8000000000,
	"black beans":            194.5000000000,
	"blue cheese":            135.0000000000,
	"blueberries":            206.5000000000,
	"bread":                  145.8000000000,
	"bread crumbs":           114.0000000000,
	"broccoli":               154.6000000000,
	"brown rice":             174.2000000000,
	"brown sugar":            160.9000000000,
	"butter":                 227.0000000000,
	"buttermilk":             245.0000000000,
	"cabbage":                112.1000000000,
	"canola oil":             180.5000000000,
	"carrot":                 155.0000000000,
	"carrots":                178.6000000000,
	"catsup":                 240.0000000000,
	"cauliflower":            142.2000000000,
	"celery":                 182.8000000000,
	"cheddar cheese":         179.0000000000,
	"cheese":                 144.7000000000,
	"cherries":               208.2000000000,
	"chicken":                194.0000000000,
	"chicken broth":          203.2000000000,
	"chicken soup":           248.3000000000,
	"chicken stock":          240.0000000000,
	"chickpeas":              193.3000000000,
	"chile":                  37.0000000000,
	"chili sauce":            259.0000000000,
	"chives":                 3.2000000000,
	"chocolate":              171.2000000000,
	"cider vinegar":          239.0000000000,
	"cinnamon":               53.5000000000,
	"cocoa":                  51.4000000000,
	"coconut":                261.9000000000,
	"coconut milk":           236.5000000000,
	"coconut oil":            218.0000000000,
	"coffee":                 248.7000000000,
	"coriander":              16.0000000000,
	"corn":                   164.1000000000,
	"corn flour":             123.4000000000,
	"cornmeal":               144.7000000000,
	"cornstarch":             128.0000000000,
	"cottage cheese":         216.3000000000,
	"cranberries":            123.3000000000,
	"cream":                  187.8000000000,
	"cream cheese":           237.3000000000,
	"cream of mushroom soup": 249.9000000000,
	"cucumber":               146.3000000000,
	"dates":                  147.0000000000,
	"dill weed":              8.9000000000,
	"dressing":               245.4000000000,
	"egg noodles":            114.2000000000,
	"eggplant":               104.0000000000,
	"evaporated milk":        252.0000000000,
	"fennel":                 87.0000000000,
	"feta cheese":            150.0000000000,
	"flour":                  113.7000000000,
	"fruit":                  225.8000000000,
	"garlic":                 144.9000000000,
	"gelatin":                195.0000000000,
	"ginger":                 96.0000000000,
	"graham crackers":        84.0000000000,
	"green beans":            220.7000000000,
	"green chilies":          241.0000000000,
	"green onions":           71.0000000000,
	"green pepper":           130.0000000000,
	"green peppers":          238.5000000000,
	"ham":                    250.4000000000,
	"hamburger":              244.0000000000,
	"hazelnuts":              108.3000000000,
	"honey":                  59.1000000000,
	"ice":                    187.3000000000,
	"kale":                   117.0000000000,
	"kidney beans":           194.6000000000,
	"leeks":                  75.0000000000,
	"lemon":                  196.9000000000,
	"lemon juice":            244.0000000000,
	"lemons":                 212.0000000000,
	"lettuce":                49.2000000000,
	"lime":                   198.0000000000,
	"lime juice":             244.0000000000,
	"mango":                  251.0000000000,
	"maple syrup":            322.0000000000,
	"margarine":              225.7000000000,
	"marshmallows":           39.3000000000,
	"mayonnaise":             230.6000000000,
	"milk":                   220.3000000000,
	"molasses":               337.0000000000,
	"mushrooms":              116.8000000000,
	"mustard":                153.0000000000,
	"nuts":                   150.6000000000,
	"oatmeal":                60.5000000000,
	"oats":                   81.0000000000,
	"oil":                    207.3000000000,
	"olive oil":              216.0000000000,
	"onion":                  152.6000000000,
	"onions":                 176.5000000000,
	"orange":                 237.7000000000,
	"orange juice":           251.2000000000,
	"oranges":                176.0000000000,
	"parmesan":               126.7000000000,
	"parmesan cheese":        100.0000000000,
	"parsley":                32.8000000000,
	"pasta":                  116.0000000000,
	"peaches":                229.9000000000,
	"peanut butter":          154.4000000000,
	"peanut oil":             216.0000000000,
	"peanuts":                135.8000000000,
	"pears":                  190.6000000000,
	"peas":                   171.1000000000,
	"pecans":                 107.0000000000,
	"pepper":                 124.0000000000,
	"pie crust":              129.0000000000,
	"pine nuts":              135.0000000000,
	"pineapple":              235.3000000000,
	"pineapple juice":        250.0000000000,
	"pork":                   146.7000000000,
	"potato":                 191.7000000000,
	"potatoes":               180.4000000000,
	"pumpkin":                134.1000000000,
	"quinoa":                 177.5000000000,
	"raisins":                140.2000000000,
	"raspberries":            179.8000000000,
	"red pepper":             125.0000000000,
	"red wine vinegar":       239.0000000000,
	"rice":                   89.2000000000,
	"ricotta cheese":         247.0000000000,
	"salad":                  207.0000000000,
	"salad oil":              214.0000000000,
	"salmon":                 177.0000000000,
	"salsa":                  250.0000000000,
	"salt":                   292.0000000000,
	"sauce":                  251.4000000000,
	"sausage":                140.3000000000,
	"scallions":              100.0000000000,
	"semolina":               107.6000000000,
	"sesame oil":             218.0000000000,
	"sesame seeds":           144.0000000000,
	"shallots":               14.4000000000,
	"shortening":             207.3000000000,
	"shrimp":                 200.2000000000,
	"skim milk":              245.0000000000,
	"soda":                   57.5000000000,
	"sour cream":             237.2000000000,
	"soy sauce":              243.5000000000,
	"spaghetti":              151.4000000000,
	"spinach":                164.1000000000,
	"spray":                  247.0000000000,
	"sugar":                  205.8000000000,
	"sweet potatoes":         224.0000000000,
	"swiss cheese":           137.7000000000,
	"tofu":                   250.0000000000,
	"tomato":                 230.5000000000,
	"tomato juice":           241.2000000000,
	"tomato sauce":           217.1000000000,
	"tomato soup":            206.2000000000,
	"tomatoes":               177.6000000000,
	"tuna":                   146.0000000000,
	"turkey":                 206.6000000000,
	"vanilla":                175.9000000000,
	"vegetable broth":        225.3000000000,
	"vegetable oil":          211.5000000000,
	"vegetable shortening":   205.0000000000,
	"vinegar":                238.0000000000,
	"walnuts":                95.0000000000,
	"water":                  243.2000000000,
	"whipped cream":          70.0000000000,
	"white bread":            41.2000000000,
	"white rice":             172.9000000000,
	"worcestershire sauce":   275.0000000000,
	"yellow cornmeal":        143.3000000000,
	"yogurt":                 211.8000000000,
	"zucchini":               194.4000000000,
}