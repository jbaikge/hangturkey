var incorrects = 0;
var animap = []
var prepAnimap = function() {
	// 1 wrong
	animap[1] = function() {
		var t = $('#Body')
		t.animate({
			height: t.height() * 1.20,
			left:   t.position().left - t.width() * 0.10,
			top:    t.position().top - t.height() * 0.10,
			width:  t.width() * 1.20
		},
		{
			duration: 300,
			queue:    true
		}).animate({
			height: t.height(),
			left:   t.position().left,
			top:    t.position().top,
			width:  t.width()
		},
		{
			duration: 100,
			queue:    true
		}).css({
			display: '',
			height:  0,
			left:    t.position().left + t.width() / 2,
			top:     t.position().top  + t.height() / 2,
			width:   0
		})
	}
	// 2 wrong
	animap[2] = function() {
		var t = $("#LeftLeg")
		t.css({
			display: '',
			top:     t.position().top - t.height()
		})
		t.animate({ top: t.position().top })
	}
	// 3 wrong
	animap[3] = function() {
		var t = $("#RightLeg")
		t.css({
			display: '',
			top:     t.position().top - t.height()
		})
		t.animate({ top: t.position().top })
	}
	// 4 wrong
	animap[4] = function() {
		var t = $('#LeftWing')
		t.css({
			display: '',
			left:    t.position().left + t.width()
		})
		t.animate({ left: t.position().left })
	}
	// 5 wrong
	animap[5] = function() {
		var t = $('#RightWing')
		t.css({
			display: '',
			left:    t.position().left - t.width()
		})
		t.animate({ left: t.position().left })
	}
	// 6 wrong
	animap[6] = function() {
		var t = $('#Neck')
		t.css({
			height: 0,
			left:   t.position().left + t.width() / 2,
			top:    t.position().top + t.height(),
			width:  0
		})
		t.animate({
			height: t.height(),
			left:   t.position().left,
			top:    t.position().top,
			width:  t.width()
		})
	}
	// 7 wrong
	animap[7] = function() {
		var rope = $('#Rope')
		var ropedTurkey = $('#Rope, #Turkey')
		rope.css({
			display: '',
			top:     -t.height()
		})
		// Drop noose
		rope.animate({
			top: t.position().top + 40
		}, {
			//queue: true
		})
		// Tighten noose around neck
		rope.animate({
			top: t.position().top
		}, {
			duration: 100//,
			//queue:    true
		})
		// Launch the turkey off stage
		ropedTurkey.animate({
			top: -$('#Stage').height()
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
	t.find('> div').css('display', 'none')
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
