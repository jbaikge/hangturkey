var incorrects = 0;
var animap = []
var prepAnimap = function(turkey) {
	var t = null
	// 0 wrong
	// animap[0] = {}
	// 1 wrong
	animap[1] = function() {
		var t = turkey.find('#Body')
		t.css({
			display: '',
			height:  0,
			left:    t.position().left + t.width() / 2,
			top:     t.position().top  + t.height() / 2,
			width:   0
		})
		t.animate({
			height: t.height() * 1.20,
			left:   t.position().left - t.width() * 0.10,
			top:    t.position().top - t.height() * 0.10,
			width:  t.width() * 1.20
		},
		{
			duration: 300,
			queue:    true
		})
		t.animate({
			height: t.height(),
			left:   t.position().left,
			top:    t.position().top,
			width:  t.width()
		},
		{
			duration: 100,
			queue:    true
		})
	}
	t = turkey.find('#Body')
	animap[1] = {
		id: t.attr('id'),
		init: {
		},
		queue: [
			{
				properties: {
					height: t.height() * 1.20,
					left:   t.position().left - t.width() * 0.10,
					top:    t.position().top - t.height() * 0.10,
					width:  t.width() * 1.20
				},
				options: {
					duration: 300,
					queue:    true
				}
			},
			{
				properties: {
					height: t.height(),
					left:   t.position().left,
					top:    t.position().top,
					width:  t.width()
				},
				options: {
					duration: 100,
					queue:    true
				}
			}
		]
	}
	// 2 wrong
	t = turkey.find('#LeftLeg')
	animap[2] = {
		id: t.attr('id'),
		init: {
			top: t.position().top - t.height()
		},
		queue: [
			{
				properties: {
					top: t.position().top
				},
				options: {
					queue: true
				}
			}
		]
	}
	// 3 wrong
	t = turkey.find('#RightLeg')
	animap[3] = {
		id: t.attr('id'),
		init: {
			top: t.position().top - t.height()
		},
		queue: [
			{
				properties: {
					top: t.position().top
				},
				options: {
					queue: true
				}
			}
		]
	}
	// 4 wrong
	t = turkey.find('#LeftWing')
	animap[4] = {
		id: t.attr('id'),
		init: {
			left: t.position().left + t.width()
		},
		queue: [
			{
				properties: {
					left: t.position().left
				},
				options: {
					queue: true
				}
			}
		]
	}
	// 5 wrong
	t = turkey.find('#RightWing')
	animap[5] = {
		id: t.attr('id'),
		init: {
			left: t.position().left - t.width()
		},
		queue: [
			{
				properties: {
					left: t.position().left
				},
				options: {
					queue: true
				}
			}
		]
	}
	// 6 wrong
	t = turkey.find('#Neck')
	animap[6] = {
		id: t.attr('id'),
		init: {
			height: 0,
			left:   t.position().left + t.width() / 2,
			top:    t.position().top + t.height(),
			width:  0
		},
		queue: [
			{
				properties: {
					height: t.height(),
					left:   t.position().left,
					top:    t.position().top,
					width:  t.width()
				},
				options: {
					queue: true
				}
			}
		]
	}
	// 7 wrong
	t = $('#Rope')
	animap[7] = {
		id: t.attr('id'),
		init: {
			top: -t.height()
		},
		queue: [
			{
				properties: {
					top: t.position().top + 40
				},
				options: {
					queue: true
				}
			},
			{
				properties: {
					top: t.position().top
				},
				options: {
					duration: 100,
					queue: true
				}
			},
			{
				target: $('#Turkey, #Rope'),
				properties: {
					top: -$('#Stage').height()
				},
				options: {
					duration: 400,
					queue: true
				}
			}
		]
	}
}
var prepDOM = function(id) {
	var t = $('<div id="Turkey"></div>')
	t.append($('<div id="LeftLeg"></div>'))
	t.append($('<div id="RightLeg"></div>'))
	t.append($('<div id="LeftWing"></div>'))
	t.append($('<div id="RightWing"></div>'))
	t.append($('<div id="Body"></div>'))
	t.append($('<div id="Neck"><div id="Head"></div></div>'))

	var r = $('<div id="Rope"></div>')

	$(id).append(t)
	$(id).append(r)
	prepAnimap(t)
	t.find('> div').css('display', 'none')
}
var addAppendage = function(idx) {
	if (animap[idx] == undefined) {
		return
	}
	var ani = animap[idx]
	var target = $('#' + ani.id)
	ani.init.display = ''
	target.css(ani.init)
	for (var i = 0; i < ani.queue.length; i++) {
		var item = ani.queue[i]
		var t = target
		if (item.target != undefined) {
			t = item.target
		}
		t.animate(item.properties, item.options)
	}
}
var letterHandler = function(e) {
	e.preventDefault()
	var self = $(this)
	var letter = self.text()
	if (letter == '') {
		$('#Error').text("Please guess a letter")
		return
	}
	if (!/[a-zA-Z]/.test(letter)) {
		$('#Error').text("You may only guess letters. Wouldn't want to expidite the turkey's doom, would you?")
		return
	}
	$.getJSON(self.attr('href'), function(data) {
		$('#TotalScore strong').text(data.TotalScore)
		$('#WordScore strong').text(data.WordScore)
		self.parent().attr('class', (data.Correct) ? 'correct' : 'wrong')
		if (data.Message != '') {
			$('#Error').text(data.Message)
			return
		}
		if (!data.Correct) {
			do {
				incorrects++
				addAppendage(incorrects)
			} while (incorrects < data.WrongCount)
		}
		$('#Word li').each(function(i, l) {
			$(l).text(data.Guessed[i])
		})
		if (data.Complete) {
			$('#NewWord').show()
		}
	})
}
$(function() {
	prepDOM('#Stage')
	$('#Alphabet a').click(letterHandler)
})
