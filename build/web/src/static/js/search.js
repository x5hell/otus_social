$(function () {

    let userInterestListTemplate = $('#userInterestListTemplate').html();
    let userCityTemplate = $('#userCityTemplate').html();
    let userTemplate = $('#userTemplate').html();
    let userListNotFoundTemplate = $('#userListNotFoundTemplate').html();

    let getUserInterestList = function(user){
        let userInterestListText = 'interestList' in user && user.interestList !== null
            ? ' / ' + user.interestList.map(function (interest) { return interest.name }).join(' / ')
            : '';
        return userInterestListTemplate
            .replace(
                '###userInterestList###',
                userInterestListText)
    };

    let getUserCity = function(user){
        return "city" in user && user.city && "id" in user.city && user.city.id > 0
            ? userCityTemplate.replace('###cityName###', user.city.name)
            : ''
    };

    let getUserProfile = function(user){
        return userTemplate
            .replace('###userId###', user.id)
            .replace('###firstName###', user.firstName)
            .replace('###lastName###', user.lastName)
            .replace('###Age###', user.age)
            .replace('###sex###', user.sex)
            .replace('###userCityTemplate###', getUserCity(user))
            .replace('###userInterestList###', getUserInterestList(user));
    };

    let showUser = function(user){
        $('#searchResult').append(getUserProfile(user));
        //console.log(getUserInterestList(user), getUserCity(user));
    };

    $('#search-form').submit(function (event) {

        $('#search-button').attr("disabled", "disabled");
        form.hideErrors();
        $.get(
            '/search',
            $('#search-form').serializeArray(),
            function (data) {
                $('#search-button').removeAttr("disabled");
                if("error" in data){
                    form.showErrors(data["error"]);
                }
                if("ok" in data){
                    $('#searchResult').empty();
                    if(data.ok != null){
                        data["ok"].forEach(showUser);
                    } else {
                        $('#searchResult').html(userListNotFoundTemplate);
                    }
                }
            },
            "json"
        );

        event.preventDefault();
    });
});