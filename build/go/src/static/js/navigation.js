$('.navbar-nav .nav-item.nav-link').each(function () {
    if( $(this).attr('href') === location.pathname){
        $(this).addClass('active');
    }
});