var incorrects = 0;
var animap = []
var prepAnimap = function() {
	// Precache element, positions, heights, and widths. As soon as display is
	// set to 'none', the positions get wacky. Also, this reduces lookups for
	// more speeeeeeeeeeed
	var parts = {}
	$('#Turkey, #Rope').children().andSelf().each(function() {
		var self = $(this)
		parts[self.attr('id')] = {
			elem:   self,
			height: self.height(),
			width:  self.width(),
			top:    self.position().top,
			left:   self.position().left
		}
	})

	// 1 wrong
	animap[1] = function() {
		// t for target, not the store
		var t = parts['Body']
		t.elem.css({
			display: '',
			height:  0,
			left:    t.left + t.width / 2,
			top:     t.top  + t.height / 2,
			width:   0
		}).animate({
			height: t.height * 1.20,
			left:   t.left - t.width * 0.10,
			top:    t.top - t.height * 0.10,
			width:  t.width * 1.20
		}, {
			duration: 300,
		}).animate({
			height: t.height,
			left:   t.left,
			top:    t.top,
			width:  t.width
		}, {
			duration: 100,
		})
	}
	// 2 wrong
	animap[2] = function() {
		var t = parts["LeftLeg"]
		t.elem.css({
			display: '',
			top:     t.top - t.height
		}).animate({ top: t.top })
	}
	// 3 wrong
	animap[3] = function() {
		var t = parts["RightLeg"]
		t.elem.css({
			display: '',
			top:     t.top - t.height
		}).animate({ top: t.top })
	}
	// 4 wrong
	animap[4] = function() {
		var t = parts['LeftWing']
		t.elem.css({
			display: '',
			left:    t.left + t.width
		}).animate({ left: t.left })
	}
	// 5 wrong
	animap[5] = function() {
		var t = parts['RightWing']
		t.elem.css({
			display: '',
			left:    t.left - t.width
		}).animate({ left: t.left })
	}
	// 6 wrong
	animap[6] = function() {
		var t = parts['Neck']
		t.elem.css({
			display: '',
			height:  0,
			left:    t.left + t.width / 2,
			top:     t.top + t.height,
			width:   0
		}).animate({
			height: t.height,
			left:   t.left,
			top:    t.top,
			width:  t.width
		})
	}
	// 7 wrong
	animap[7] = function() {
		var rope = parts["Rope"]
		var ropedTurkey = parts["Rope"].elem.add(parts["Turkey"].elem)
		var stage = $('#Stage')
		rope.elem.css({
			display: '',
			top:     -rope.height
		})
		stage.queue(function(next) {
			// Drop noose
			rope.elem.animate({
				top: rope.top + 40
			}, next)
		}).queue(function(next) {
			// Pause for drama
			setTimeout(next, 250)
		}).queue(function(next) {
			// Tighten noose around neck
			rope.elem.animate({
				top: rope.top
			}, 75, next)
		}).queue(function(next) {
			// Enlarge eyeballs
			$('#Head').addClass('alarmed')
			next()
		}).queue(function(next) {
			// Pause for drama
			setTimeout(next, 500)
		}).queue(function(next) {
			// Launch the turkey off stage
			ropedTurkey.animate({
				top: stage.height() * -1 // FireFox loses it's balls with the leading -
			}, next)
		})
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
	prepAnimap()
	t.find('> div').add(r).css('display', 'none')
}
var addAppendage = function(idx) {
	if (animap[idx] == undefined) {
		return
	}
	animap[idx]()
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
