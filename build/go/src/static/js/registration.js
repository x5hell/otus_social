$(function () {
    $('#registration-form').submit(function (event) {

        $('#registration-button').attr("disabled", "disabled");
        $('.validation-element').each(function () {
            hideError($(this).attr('id'));
        });
        $.post(
            '/registration',
            $('#registration-form').serializeArray(),
            function (data) {
                $('#registration-button').removeAttr("disabled");
                if("error" in data){
                    showErrors(data["error"]);
                }
                if("ok" in data){

                }
            },
            "json"
        );

        event.preventDefault();
    });

    $('.validation-element').focusin(function () {
        hideError($(this).attr('id'));
    });

    let showErrors = function(errorList) {
        for(let fieldId in errorList){
            let errorMessage = errorList[fieldId];
            showError(fieldId, errorMessage);
        }
        $([document.documentElement, document.body]).animate({
            scrollTop: $('.is-invalid').first().offset().top
        }, 500);
    };

    let showError = function(fieldId, errorMessage){
        $('#' + fieldId)
            .addClass('form-control is-invalid')
            .popover({
                placement: 'right',
                container: 'body',
                content: '<i class="text-danger">' + errorMessage + '</i>',
                trigger: 'manual',
                html: true
            });
        $('#' + fieldId)
            .popover("show");
    };

    let hideError = function(fieldId){
        $('#' + fieldId)
            .removeClass('form-control is-invalid')
            .popover('dispose')
    };
});