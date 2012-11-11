// jQuery form oninput wrapper
// Thx to Mathias Bynens https://github.com/mathiasbynens/jquery-oninput
jQuery.fn.input=function(a){var b=this;return a?b.bind({'input.input':function(c){b.unbind('keydown.input');a.call(this,c)},'keydown.input':function(c){a.call(this,c)}}):b.trigger('keydown.input')};

$(document).ready(function() {

	function testRegex () {
		$.ajax({
			type: "POST",
  			url: "/test_regexp/",
			data: {
				regexp: $("#regexpInput").val()
				, testString: $("#testStringInput").val()
				, findAll: $("#findAllCheckbox").is(":checked")
			}
		}).done(function(msg) {
			var res = $.parseJSON(msg);
			var allMatches = res.matches;
			var groupsName = res.groupsName;
			console.log(res);

			// Empty match groups table
			$("#matchGroupsTable tbody > tr").remove();

			if (allMatches && allMatches.length > 0)
			{
				var l = allMatches.length;

				// Match font color > green
				$("#match").css("color", "green")

				// Match groups
				if (l > 0)
				{
					var match = [];
					var matchGroupsTable = [];
					var index = 0;

					for (var i = 0; i < l; i++)
					{
						var matches = allMatches[i]; 
						var m = matches.length;

						match.push(matches[0]);
						for (var j = 1; j < m; j++) 
						{
							matchGroupsTable.push('<tr><td>'+(index++)+'</td><td>'+((groupsName[j-1] != "") ? groupsName[j-1] : "-")+'</td><td>'+matches[j]+'</td></tr>');
						}
					}

					$("#match").html(match.join(" "))
					$('#matchGroupsTable > tbody:last').append(matchGroupsTable.join());
				}

			}
			else
			{
				$("#match").html("No match !")
				$("#match").css("color", "red")
			}

		});
	}

	$("#regexpForm").submit(function() {
		testRegex();

		return false;
	});

	var testRegexTimeout;

	$('#regexpForm').input(function() {
		
		if (testRegexTimeout)
			clearTimeout(testRegexTimeout)

		testRegexTimeout = setTimeout(testRegex, 750)
	});

	testRegex();

});