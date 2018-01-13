$(function() {
    var url = getBaseUrl()[0];
    displayLed();

    $('#submitMessage').click(function(){
        var message = {
            'text': $('#message').val(),
            'duration' : 10
        }

        $.ajax({
            type: "POST",
            url: url+"api/message",
            data: JSON.stringify(message),
            success: function(){
                alert('Sent!')   
            },
            fail: function() {
                alert( "error" );
            }
        });
    })

    $('#submitForecast').click(function(){
        $.ajax({
            type: "POST",
            url: url+"api/forecast",
            success: function(){
                alert('Sent!')   
            },
            fail: function() {
                alert( "error" );
            }
        });
    })

    $('#submitWeather').click(function(){
        var message = {
            'text': $('#message').val(),
            'duration' : 10
        }

        $.ajax({
            type: "POST",
            url: url+"api/weather",
            success: function(){
                alert('Sent!')   
            },
            fail: function() {
                alert( "error" );
            }
        });
    })

    $('#submitLED').click(function(){
        $.ajax({
            type: "POST",
            url: url+"api/led",
            success: function(){
                alert('Sent!')   
            },
            fail: function() {
                alert( "error" );
            }
        });

        displayLed();
    })
    
    function getBaseUrl() {
        var re = new RegExp(/^.*\//);
        return re.exec(window.location.href);
    }

    function displayLed() {
        $.ajax({
            type: "GET",
            url: url+"api/status",
            success: function(data){
                $('#ledActive').text("Is LED Active: "+data.isLEDActive); 
                if(data.isLEDActive) {
                    $('#status').removeClass("error");
                    $('#status').addClass("ok");
                }
                else {
                    $('#status').removeClass("ok");
                    $('#status').addClass("error");
                }
            },
            fail: function() {
                $('#ledActive').text("error");
                $('#status').removeClass("ok");
                $('#status').addClass("error");
            }
        });
    }
});