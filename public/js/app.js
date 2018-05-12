/**
 * Created by Tanvir on 2018-04-27.
 */

$(function () {

    // For registration or join
    var $fname = $("#firstname");
    var $lname = $("#lastname");
    var $email = $("#email");
    var $password = $("#password");
    var $contactno = $("#contactno");
    var $vjudge = $("#vjudge");

    $("#btn-sign-up").on('click', function () {
        var user = {
            firstname: $fname.val(),
            lastname: $lname.val(),
            email: $email.val(),
            password: $password.val(),
            contactno: $contactno.val(),
            vjudge: $vjudge.val(),
        };

        $.ajax({
            url: '/join',
            type: 'POST',
            data: JSON.stringify(user),
            dataType: "json",
            contentType: "application/json",
            success: function (response) {
                // redirect must be defined and must be true
                if (response.redirect !== undefined && response.redirect) {
                    window.location.href = response.redirectUrl;
                }
            }
        });
    });

    // For Login

    var $email_in = $("#email_in");
    var $password_in = $("#password_in");
    $("#btn-sign-in").on('click', function () {
        var user = {
            email: $email_in.val(),
            password: $password_in.val()
        };
        
        $.ajax({
            url: '/login',
            type: 'POST',
            data: JSON.stringify(user),
            dataType: "json",
            contentType: "application/json",
            success: function (response) {
                // redirect must be defined and must be true
                if (response.redirect !== undefined && response.redirect) {
                    window.location.href = response.redirectUrl;
                }
            }
        });
    });


});


$(function () {
    // Handle blog create button
    var $tittle = $("#tittle");
    var $description = $("#description");

    $("#btn-post").on('click', function () {
        var pst = {
            tittle: $tittle.val(),
            description: $description.val(),
        };

        $.ajax({
            url: '/blog/create',
            type: 'POST',
            data: JSON.stringify(pst),
            dataType: "json",
            contentType: "application/json",
            success: function (response) {
                // redirect must be defined and must be true
                if (response.redirect !== undefined && response.redirect) {
                    window.location.href = response.redirectUrl;
                }
            }
        });
    });
});


$(function () {
    // For Contest Create fab button
    $( "#cCreateButton" ).click(function(){
        $( "#cCreateDialog" ).dialog("open")
    });

    $( "#cCreateDialog" ).dialog({
        draggable:false,
        resizable:false,
        closeOnEscape:false,
        modal:true,
        autoOpen:false,
        minWidth:585
    });
    $(".ui-dialog-titlebar").hide();

    $( "#cCreateClear" ).click(function(){
        $( "#cCreateDialog" ).dialog("close")
    });
});