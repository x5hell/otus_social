$(function () {
    $('#login-form').submit(function (event) {

        $('#login-button').attr("disabled", "disabled");
        $('.validation-element').each(function () {
            form.hideError($(this).attr('id'));
        });
        $.post(
            '/login',
            $('#login-form').serializeArray(),
            function (data) {
                $('#login-button').removeAttr("disabled");
                if("error" in data){
                    form.showErrors(data["error"]);
                }
                if("ok" in data){
                    window.location.href = "/edit-profile-form"
                }
            },
            "json"
        );

        event.preventDefault();
    });
});