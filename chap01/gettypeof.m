% gettypeof Return the type of an object
% gettypeof(variable, whos)
% 
% a = 1;
% b = 'abc';
% c = linspace(1, 100);
% whos
%   Name      Size             Bytes  Class     Attributes
% 
%   a         1x1                  8  double              
%   b         1x3                  6  char                
%   c         1x100              800  double        
% 
% >> fprintf("%s\n", gettypeof('b', whos));
% char
function ret_size = gettypeof(var_name, ret_whos)
    ret_size = -1;

    for idx = 1 : length(ret_whos)
        if ret_whos(idx).name == var_name
            ret_size = ret_whos(idx).class;
        end
    end

    if ret_size == -1
        error("Invalid variable name: %s", var_name);
    end
end