import React from 'react';
import PropTypes from 'prop-types';

import Menu from 'Components/Menu';
import Tooltip from 'Components/Tooltip';
import TooltipOverlay from 'Components/TooltipOverlay';

const RowActionMenu = ({
    text,
    icon,
    border,
    className,
    menuClassName,
    buttonClassName,
    options,
    dataTestId
}) => (
    <Tooltip content={<TooltipOverlay>{text}</TooltipOverlay>}>
        <div>
            <Menu
                className={`${className} ${border}`}
                menuClassName={menuClassName}
                buttonClass={`p-1 px-4 ${buttonClassName}`}
                buttonContent={icon}
                options={options}
                dataTestId={dataTestId}
            />
        </div>
    </Tooltip>
);

RowActionMenu.propTypes = {
    text: PropTypes.string.isRequired,
    icon: PropTypes.node.isRequired,
    border: PropTypes.string,
    className: PropTypes.string,
    menuClassName: PropTypes.string,
    buttonClassName: PropTypes.string,
    options: PropTypes.oneOfType([
        PropTypes.arrayOf(
            PropTypes.shape({
                className: PropTypes.string,
                icon: PropTypes.object,
                label: PropTypes.string.isRequired,
                link: PropTypes.string,
                onClick: PropTypes.func
            })
        ).isRequired,
        PropTypes.shape({})
    ]).isRequired,
    dataTestId: PropTypes.string
};

RowActionMenu.defaultProps = {
    className: 'hover:bg-primary-200 text-primary-600 hover:text-primary-700',
    menuClassName: 'bg-base-200 min-w-28',
    buttonClassName: 'hover:bg-primary-200 text-primary-600 hover:text-primary-700',
    border: '',
    dataTestId: ''
};

export default RowActionMenu;
