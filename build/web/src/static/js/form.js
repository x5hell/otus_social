function Form() {}

Form.prototype = {
    showErrors: function(errorList) {
        for(let fieldId in errorList){
            let errorMessage = errorList[fieldId];
            this.showError(fieldId, errorMessage);
        }
        $([document.documentElement, document.body]).animate({
            scrollTop: $('.is-invalid').first().offset().top
        }, 500);
    },

    showError: function(fieldId, errorMessage){
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
    },

    hideErrors: function(){
        let self = this;
        $('.validation-element').each(function () {
            self.hideError($(this).attr('id'));
        });
    },

    hideError: function(fieldId){
        $('#' + fieldId)
            .removeClass('form-control is-invalid')
            .popover('dispose')
    }
};

$(function(){
    window.form = new Form();
    $('.validation-element').focusin(function () {
        form.hideError($(this).attr('id'));
    });
});