function v_pi = calculate_pi(n_term)
    numerator = 4.0;
    denominator = 1.0;
    operation = 1.0;
    v_pi = 0.0;

    for i = 0 : n_term
        v_pi = v_pi + operation * (numerator / denominator);
        denominator = denominator + 2.0;
        operation = operation * (- 1.0);
    end
end