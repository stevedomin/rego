// jQuery form oninput wrapper
// Thx to Mathias Bynens https://github.com/mathiasbynens/jquery-oninput
jQuery.fn.input=function(a){var b=this;return a?b.bind({'input.input':function(c){b.unbind('keydown.input');a.call(this,c)},'keydown.input':function(c){a.call(this,c)}}):b.trigger('keydown.input')};

$(document).ready(function() {

	function testRegex () {
		$.ajax({
			type: "POST",
  			url: "/test_regex/",
			data: {
				regex: $("#regexInput").val()
				, testString: $("#testStringInput").val()
			}
		}).done(function(msg) {
			var res = $.parseJSON(msg);

			console.log(res);

			// Empty match groups table
			$("#matchGroupsTable tbody > tr").remove();

			if (res.match == "")
			{
				$("#match").html("No match !")
				$("#match").css("color", "red")
			}
			else
			{
				var l = res.groups.length;

				// Match
				$("#match").html(res.match)
				$("#match").css("color", "green")

				// Match groups
				if (l > 0)
				{
					for (var i = 0; i < l; i++)
					{
						$('#matchGroupsTable > tbody:last').append('<tr><td>'+i+'</td><td>'+((res.groupsName[i] != "") ? res.groupsName[i] : "-")+'</td><td>'+res.groups[i]+'</td></tr>');
					}
				}

			}

		});
	}

	$("#regexForm").submit(function() {
		testRegex();

		return false;
	});

	var testRegexTimeout;

	$('#regexForm').input(function() {
		
		if (testRegexTimeout)
			clearTimeout(testRegexTimeout)

		testRegexTimeout = setTimeout(testRegex, 750)
	});

});