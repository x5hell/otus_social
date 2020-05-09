$(function () {
    $('#edit-profile-form').submit(function (event) {

        $('#edit-profile-button').attr("disabled", "disabled");
        form.hideErrors();
        $.post(
            '/edit-profile',
            $('#edit-profile-form').serializeArray(),
            function (data) {
                $('#edit-profile-button').removeAttr("disabled");
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