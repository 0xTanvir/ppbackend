/**
 * Created by Tanvir on 2018-04-27.
 */

 $(function (){
     var $fname = $("#firstname");
     var $lname = $("#lastname");
     var $email = $("#email");
     var $password = $("#password");
     var $contactno = $("#contactno");
     var $vjudge = $("#vjudge");

     $("#btn-sign-up").on('click', function() {
        var user = {
            firstname : $fname.val(),
            lastname : $lname.val(),
            email : $email.val(),
            password : $password.val(),
            contactno : $contactno.val(),
            vjudge : $vjudge.val(),
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
 });
